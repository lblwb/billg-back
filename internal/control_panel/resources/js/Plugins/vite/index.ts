// This code is taken from laravel vite plugin: https://github.com/laravel/vite-plugin/blob/1.x/src/index.ts

import fs from 'fs'
import {AddressInfo} from 'net'
import path from 'path'
import colors from 'picocolors'
import {ConfigEnv, loadEnv, Plugin, PluginOption, ResolvedConfig, SSROptions, UserConfig} from 'vite'
import fullReload, {Config as FullReloadConfig} from 'vite-plugin-full-reload'

interface PluginConfig {
    /**
     * The path or paths of the entry points to compile.
     */
    input: string | string[]

    /**
     * Artefak's public directory.
     *
     * @default 'public'
     */
    publicDirectory?: string

    /**
     * The public subdirectory where compiled assets should be written.
     *
     * @default 'build'
     */
    buildDirectory?: string

    /**
     * The path to the "hot" file.
     *
     * @default `${publicDirectory}/hot`
     */
    hotFile?: string

    /**
     * The path of the SSR entry point.
     */
    ssr?: string | string[]

    /**
     * The directory where the SSR bundle should be written.
     *
     * @default 'bootstrap/ssr'
     */
    ssrOutputDirectory?: string

    /**
     * Configuration for performing full page refresh on blade (or other) file changes.
     *
     * {@link https://github.com/ElMassimo/vite-plugin-full-reload}
     * @default false
     */
    refresh?: boolean | string | string[] | RefreshConfig | RefreshConfig[]

    /**
     * Transform the code while serving.
     */
    transformOnServe?: (code: string, url: DevServerUrl) => string,
}

interface RefreshConfig {
    paths: string[],
    config?: FullReloadConfig,
}

interface ArtefakPlugin extends Plugin {
    config: (config: UserConfig, env: ConfigEnv) => UserConfig
}

type DevServerUrl = `${'http' | 'https'}://${string}:${number}`

let exitHandlersBound = false

export const refreshPaths = [
    'resources/views/**',
    'routes/**',
]

/**
 * Artefak plugin for Vite.
 *
 * @param config - A config object or relative path(s) of the scripts to be compiled.
 */
export default function artefak(config: string | string[] | PluginConfig): [ArtefakPlugin, ...Plugin[]] {
    const pluginConfig = resolvePluginConfig(config)

    return [
        resolveArtefakPlugin(pluginConfig),
        ...resolveFullReloadConfig(pluginConfig) as Plugin[],
    ];
}

/**
 * Resolve the Artefak Plugin configuration.
 */
function resolveArtefakPlugin(pluginConfig: Required<PluginConfig>): ArtefakPlugin {
    let viteDevServerUrl: DevServerUrl
    let resolvedConfig: ResolvedConfig
    let userConfig: UserConfig

    const defaultAliases: Record<string, string> = {
        '@': path.resolve(__dirname, '../../../../resources/js'),
    };

    return {
        name: 'artefak',
        enforce: 'post',
        config: (config, {command, mode}) => {
            userConfig = config
            const ssr = !!userConfig.build?.ssr
            const env = loadEnv(mode, userConfig.envDir || process.cwd(), '')
            const assetUrl = env.ASSET_URL ?? ''
            // @ts-ignore
            const serverConfig = command === 'air'
                ? resolveEnvironmentServerConfig(env)
                : undefined

            ensureCommandShouldRunInEnvironment(command, env)

            return {
                base: userConfig.base ?? (command === 'build' ? resolveBase(pluginConfig, assetUrl) : ''),
                publicDir: userConfig.publicDir ?? false,
                build: {
                    manifest: userConfig.build?.manifest ?? (ssr ? false : 'manifest.json'),
                    ssrManifest: userConfig.build?.ssrManifest ?? (ssr ? 'ssr-manifest.json' : false),
                    outDir: userConfig.build?.outDir ?? resolveOutDir(pluginConfig, ssr),
                    rollupOptions: {
                        input: userConfig.build?.rollupOptions?.input ?? resolveInput(pluginConfig, ssr)
                    },
                    assetsInlineLimit: userConfig.build?.assetsInlineLimit ?? 0,
                },
                server: {
                    origin: userConfig.server?.origin ?? '__artefak_vite_placeholder__',
                    ...(serverConfig ? {
                        host: userConfig.server?.host ?? serverConfig.host,
                        hmr: userConfig.server?.hmr === false ? false : {
                            ...serverConfig.hmr,
                            ...(userConfig.server?.hmr === true ? {} : userConfig.server?.hmr),
                        },
                        https: userConfig.server?.https ?? serverConfig.https,
                    } : undefined),
                },
                resolve: {
                    alias: Array.isArray(userConfig.resolve?.alias)
                        ? [
                            ...userConfig.resolve?.alias ?? [],
                            ...Object.keys(defaultAliases).map(alias => ({
                                find: alias,
                                replacement: defaultAliases[alias]
                            }))
                        ]
                        : {
                            ...defaultAliases,
                            ...userConfig.resolve?.alias,
                        }
                },
                ssr: {
                    noExternal: noExternalInertiaHelpers(userConfig),
                },
            }
        },
        configResolved(config) {
            resolvedConfig = config
        },
        transform(code) {
            // @ts-ignore
            if (resolvedConfig.command === 'air') {
                code = code.replace(/__artefak_vite_placeholder__/g, viteDevServerUrl)

                return pluginConfig.transformOnServe(code, viteDevServerUrl)
            }
        },
        configureServer(server) {
            const envDir = resolvedConfig.envDir || process.cwd()
            const appUrl = loadEnv(resolvedConfig.mode, envDir, 'APP_URL').APP_URL ?? 'undefined'

            server.httpServer?.once('listening', () => {
                const address = server.httpServer?.address()

                const isAddressInfo = (x: string | AddressInfo | null | undefined): x is AddressInfo => typeof x === 'object'
                if (isAddressInfo(address)) {
                    viteDevServerUrl = userConfig.server?.origin ? userConfig.server.origin as DevServerUrl : resolveDevServerUrl(address, server.config, userConfig)
                    fs.writeFileSync(pluginConfig.hotFile, viteDevServerUrl)

                    setTimeout(() => {
                        server.config.logger.info(`\n  ${colors.red(`${colors.bold('Artefak')} ${artefakVersion()}`)}  ${colors.dim('plugin')} ${colors.bold(`v${pluginVersion()}`)}`)
                        server.config.logger.info('')
                        server.config.logger.info(`  ${colors.green('➜')}  ${colors.bold('APP_URL')}: ${colors.cyan(appUrl.replace(/:(\d+)/, (_, port) => `:${colors.bold(port)}`))}`)
                    }, 100)
                }
            })

            if (!exitHandlersBound) {
                const clean = () => {
                    if (fs.existsSync(pluginConfig.hotFile)) {
                        fs.rmSync(pluginConfig.hotFile)
                    }
                }

                process.on('exit', clean)
                process.on('SIGINT', () => process.exit())
                process.on('SIGTERM', () => process.exit())
                process.on('SIGHUP', () => process.exit())

                exitHandlersBound = true
            }

            return () => server.middlewares.use((req, res, next) => {
                next()
            })
        }
    }
}

/**
 * Validate the command can run in the given environment.
 */
function ensureCommandShouldRunInEnvironment(command: "build" | "serve", env: Record<string, string>): void {
    if (command === 'build' || env.ARTEFAK_BYPASS_ENV_CHECK === '1') {
        return;
    }

    if (typeof env.CI !== 'undefined') {
        throw Error('You should not run the Vite HMR server in CI environments. You should build your assets for production instead. To disable this ENV check you may set ARTEFAK_BYPASS_ENV_CHECK=1')
    }
}

/**
 * The version of Artefak being run.
 */
function artefakVersion(): string {
    return '1.0.0'
}

/**
 * The version of the Artefak Vite plugin being run.
 */
function pluginVersion(): string {
    return '1.0.0'
}

/**
 * Convert the users configuration into a standard structure with defaults.
 */
function resolvePluginConfig(config: string | string[] | PluginConfig): Required<PluginConfig> {
    if (typeof config === 'undefined') {
        throw new Error('artefak-vite-plugin: missing configuration.')
    }

    if (typeof config === 'string' || Array.isArray(config)) {
        config = {input: config, ssr: config}
    }

    if (typeof config.input === 'undefined') {
        throw new Error('artefak-vite-plugin: missing configuration for "input".')
    }

    if (typeof config.publicDirectory === 'string') {
        config.publicDirectory = config.publicDirectory.trim().replace(/^\/+/, '')

        if (config.publicDirectory === '') {
            throw new Error('artefak-vite-plugin: publicDirectory must be a subdirectory. E.g. \'public\'.')
        }
    }

    if (typeof config.buildDirectory === 'string') {
        config.buildDirectory = config.buildDirectory.trim().replace(/^\/+/, '').replace(/\/+$/, '')

        if (config.buildDirectory === '') {
            throw new Error('artefak-vite-plugin: buildDirectory must be a subdirectory. E.g. \'build\'.')
        }
    }

    if (typeof config.ssrOutputDirectory === 'string') {
        config.ssrOutputDirectory = config.ssrOutputDirectory.trim().replace(/^\/+/, '').replace(/\/+$/, '')
    }

    if (config.refresh === true) {
        config.refresh = [{paths: refreshPaths}]
    }

    return {
        input: config.input,
        publicDirectory: config.publicDirectory ?? 'public',
        buildDirectory: config.buildDirectory ?? 'build',
        ssr: config.ssr ?? config.input,
        ssrOutputDirectory: config.ssrOutputDirectory ?? 'bootstrap/ssr',
        refresh: config.refresh ?? false,
        hotFile:  config.hotFile ?? path.join((config.publicDirectory ?? 'public'), 'hot'),
        transformOnServe: config.transformOnServe ?? ((code) => code),
    }
}

/**
 * Resolve the Vite base option from the configuration.
 */
function resolveBase(config: Required<PluginConfig>, assetUrl: string): string {
    return assetUrl + (!assetUrl.endsWith('/') ? '/' : '') + config.buildDirectory + '/'
}

/**
 * Resolve the Vite input path from the configuration.
 */
function resolveInput(config: Required<PluginConfig>, ssr: boolean): string | string[] | undefined {
    if (ssr) {
        return config.ssr
    }

    return config.input
}

/**
 * Resolve the Vite outDir path from the configuration.
 */
function resolveOutDir(config: Required<PluginConfig>, ssr: boolean): string | undefined {
    if (ssr) {
        return config.ssrOutputDirectory
    }

    return path.join(config.publicDirectory, config.buildDirectory)
}

function resolveFullReloadConfig({refresh: config}: Required<PluginConfig>): PluginOption[] {
    if (typeof config === 'boolean') {
        return [];
    }

    if (typeof config === 'string') {
        config = [{paths: [config]}]
    }

    if (!Array.isArray(config)) {
        config = [config]
    }

    if (config.some(c => typeof c === 'string')) {
        config = [{paths: config}] as RefreshConfig[]
    }

    return (config as RefreshConfig[]).flatMap(c => {
        const plugin = fullReload(c.paths, c.config)

        /* eslint-disable-next-line @typescript-eslint/ban-ts-comment */
        /** @ts-ignore */
        plugin.__artefak_plugin_config = c

        return plugin
    })
}

/**
 * Resolve the dev server URL from the server address and configuration.
 */
function resolveDevServerUrl(address: AddressInfo, config: ResolvedConfig, userConfig: UserConfig): DevServerUrl {
    const configHmrProtocol = typeof config.server.hmr === 'object' ? config.server.hmr.protocol : null
    const clientProtocol = configHmrProtocol ? (configHmrProtocol === 'wss' ? 'https' : 'http') : null
    const serverProtocol = config.server.https ? 'https' : 'http'
    const protocol = clientProtocol ?? serverProtocol

    const configHmrHost = typeof config.server.hmr === 'object' ? config.server.hmr.host : null
    const configHost = typeof config.server.host === 'string' ? config.server.host : null
    const serverAddress = isIpv6(address) ? `[${address.address}]` : address.address
    const host = configHmrHost ?? configHost ?? serverAddress

    const configHmrClientPort = typeof config.server.hmr === 'object' ? config.server.hmr.clientPort : null
    const port = configHmrClientPort ?? address.port

    return `${protocol}://${host}:${port}`
}

function isIpv6(address: AddressInfo): boolean {
    return address.family === 'IPv6'
        // In node >=18.0 <18.4 this was an integer value. This was changed in a minor version.
        // See: https://github.com/laravel/vite-plugin/issues/103
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore-next-line
        || address.family === 6;
}

/**
 * Add the Inertia helpers to the list of SSR dependencies that aren't externalized.
 *
 * @see https://vitejs.dev/guide/ssr.html#ssr-externals
 */
function noExternalInertiaHelpers(config: UserConfig): true | Array<string | RegExp> {
    /* eslint-disable-next-line @typescript-eslint/ban-ts-comment */
    /* @ts-ignore */
    const userNoExternal = (config.ssr as SSROptions | undefined)?.noExternal
    const pluginNoExternal = ['artefak-vite-plugin']

    if (userNoExternal === true) {
        return true
    }

    if (typeof userNoExternal === 'undefined') {
        return pluginNoExternal
    }

    return [
        ...(Array.isArray(userNoExternal) ? userNoExternal : [userNoExternal]),
        ...pluginNoExternal,
    ]
}

/**
 * Resolve the server config from the environment.
 */
function resolveEnvironmentServerConfig(env: Record<string, string>): {
    hmr?: { host: string }
    host?: string,
    https?: { cert: Buffer, key: Buffer }
} | undefined {
    if (!env.VITE_DEV_SERVER_KEY && !env.VITE_DEV_SERVER_CERT) {
        return
    }

    if (!fs.existsSync(env.VITE_DEV_SERVER_KEY) || !fs.existsSync(env.VITE_DEV_SERVER_CERT)) {
        throw Error(`Unable to find the certificate files specified in your environment. Ensure you have correctly configured VITE_DEV_SERVER_KEY: [${env.VITE_DEV_SERVER_KEY}] and VITE_DEV_SERVER_CERT: [${env.VITE_DEV_SERVER_CERT}].`)
    }

    const host = resolveHostFromEnv(env)

    if (!host) {
        throw Error(`Unable to determine the host from the environment's APP_URL: [${env.APP_URL}].`)
    }

    return {
        hmr: {host},
        host,
        https: {
            key: fs.readFileSync(env.VITE_DEV_SERVER_KEY),
            cert: fs.readFileSync(env.VITE_DEV_SERVER_CERT),
        },
    }
}

/**
 * Resolve the host name from the environment.
 */
function resolveHostFromEnv(env: Record<string, string>): string | undefined {
    try {
        return new URL(env.APP_URL).host
    } catch {
        return
    }
}

