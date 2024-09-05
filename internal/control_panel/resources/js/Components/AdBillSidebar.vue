<script>
import {defineComponent} from 'vue'
import {router} from "@inertiajs/vue3";

export default defineComponent({
    name: "AdBillSidebar",
    data() {
        return {
            billState: {
                services: [{}],
            },
            client_system: {
                client: {
                    count: {
                        new_count: 100,
                        active_count: 150,
                    }
                },
            },
            order_system: {
                client: {
                    count: {
                        new_count: 20,
                        active_count: 150,
                    }
                },
            },
            supp_system: {
                client: {
                    count: {
                        new_count: 4,
                        active_count: 150,
                    }
                },
            }
        }
    },
    methods: {
        router() {
            return router
        }
    }

})
</script>
<template>
    <div class="Sidebar">
        <div class="SidebarWrapper">
            <div class="SidebarLogo" style="cursor:pointer;">
                <div class="SidebarLogoWrapper">
                    <a @click.prevent="router().get('/')">
                        <!--                        <h2 style="color: var(&#45;&#45;primaryColor)!important;">Управление Billing</h2>-->

                    </a>
                </div>
            </div>
            <div class="SidebarNavbar">
                <!-- Клиенты -->
                <div class="SidebarNavbarChap" v-if="billState.services && billState.services.length > 0">
                    <div class="SidebarNavbarHeading">
                        <div class="SidebarNavbarHeadingTitle">
                            Управление клиентами
                        </div>
                    </div>
                    <div class="SidebarNavbarWrapper">
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/panel/bill/clients/all')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>👥 Клиенты</span>
                                        <span style="color: blueviolet"> [n-{{
                                            client_system.client.count.new_count
                                            }}]</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/panel/bill/orders/all')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span style="color: red">🛍️ Заказы</span> <span
                                            style="color: red"> [n-{{
                                        order_system.client.count.new_count
                                        }}]</span>
                                    </div>
                                </div>
                            </a>
                        </div>

<!--                        <div class="SidebarNavbarItem">-->
<!--                            <a @click="() => {alert('Пока-что не доступы!')}"-->
<!--                               class="SidebarNavbarItemLink">-->
<!--                                <div class="SidebarNavbarItemLinkHeading">-->
<!--                                    <div class="SidebarNavbarItemLinkHeadingTitle">-->
<!--                                        Счета клиентов-->
<!--                                    </div>-->
<!--                                </div>-->
<!--                            </a>-->
<!--                        </div>-->
<!--                        <div class="SidebarNavbarItem">-->
<!--                            <a @click="() => {alert('Пока-что не доступы!')}"-->
<!--                               class="SidebarNavbarItemLink">-->
<!--                                <div class="SidebarNavbarItemLinkHeading">-->
<!--                                    <div class="SidebarNavbarItemLinkHeadingTitle">-->
<!--                                        <span style="color: red">Поддержка клиентов</span> <span-->
<!--                                            style="color: red">[n-{{-->
<!--                                        supp_system.client.count.new_count-->
<!--                                        }}]</span>-->
<!--                                    </div>-->
<!--                                </div>-->
<!--                            </a>-->
<!--                        </div>-->
                    </div>
                </div>

                <!-- Продукты -->
                <div class="SidebarNavbarChap" v-if="billState.services && billState.services.length > 0">
                    <div class="SidebarNavbarHeading">
                        <div class="SidebarNavbarHeadingTitle">
                            Управление наличием продуктов
                        </div>
                    </div>
                    <div class="SidebarNavbarWrapper">
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/panel/bill/services/all')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>🛎️ Услуги</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                        <!--                            <a @click="router().get('/panel/bill/services/tariffs/all')"-->
<!--                        <div class="SidebarNavbarItem">-->
<!--                           -->
<!--                            <a @click="() => {alert('Пока-что не доступы!')}"-->
<!--                               class="SidebarNavbarItemLink">-->
<!--                                <div class="SidebarNavbarItemLinkHeading">-->
<!--                                    <div class="SidebarNavbarItemLinkHeadingTitle">-->
<!--                                        <span>Тарифы услуг</span>-->
<!--                                    </div>-->
<!--                                </div>-->
<!--                            </a>-->
<!--                        </div>-->
                    </div>
                </div>

                <div class="SidebarNavbarChap" v-if="billState.services && billState.services.length > 0">
                    <div class="SidebarNavbarHeading">
                        <div class="SidebarNavbarHeadingTitle">
                            Управление рассылками
                        </div>
                    </div>
                    <div class="SidebarNavbarWrapper">
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>📰 Листинг новостей</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>📟 Telegram-рассылка</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>📻 Telegram-канал</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                        <div class="SidebarNavbarItem">
                            <a @click="router().get('/')"
                               class="SidebarNavbarItemLink">
                                <div class="SidebarNavbarItemLinkHeading">
                                    <div class="SidebarNavbarItemLinkHeadingTitle">
                                        <span>🌐 Социальные-сети</span>
                                    </div>
                                </div>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>


.Sidebar {
    min-width: 320px;
    max-width: 320px;
    height: 100vh;
    max-height: 100%;
    flex-shrink: 0;
    background: #111;
    border-right: solid 2px #222;
    position: fixed;
    top: 0;
    left: 0;
}

.Sidebar .SidebarWrapper {
    padding: 24px 36px;
    height: 100%;
}

.Sidebar .SidebarLogo {
    margin-bottom: 36px;
}

.Sidebar .SidebarNavbarHeading {
    margin-bottom: 24px;
}

.Sidebar .SidebarNavbarHeading .SidebarNavbarHeadingTitle {
    color: #3a3939;
    font-family: 'Inter', sans-serif;
    font-size: 12px;
    font-style: normal;
    font-weight: 500;
    line-height: normal;
}

.SidebarNavbarItem {
    margin-bottom: 16px;
    padding: 10px;
    border: solid 1px #191919;
    border-radius: 16px;
}

.SidebarNavbarItem:hover,
.SidebarNavbarItem:focus {
    border: solid 1px #444;
}

.SidebarNavbarItem .SidebarNavbarItemLink {
    text-decoration: none;
}

.SidebarNavbarItem .SidebarNavbarItemLink .SidebarNavbarItemLinkHeading {

}

.SidebarNavbarItem .SidebarNavbarItemLink .SidebarNavbarItemLinkHeading {

}

.SidebarNavbar .SidebarNavbarHeading .SidebarNavbarHeadingTitle {
    font-size: 12px;
    font-weight: 600;
    text-transform: uppercase;
    color: #555;
}

.SidebarNavbarItem .SidebarNavbarItemLink .SidebarNavbarItemLinkHeading .SidebarNavbarItemLinkHeadingTitle {
    color: #8B8585;
    font-family: 'Inter', sans-serif;
    font-size: 16px;
    font-style: normal;
    font-weight: 500;
    line-height: normal;
    cursor: pointer;
}

/*Active*/
.SidebarNavbarItem .SidebarNavbarItemLink.Active .SidebarNavbarItemLinkHeading .SidebarNavbarItemLinkHeadingTitle {
    color: #0B4F3C;
}

/*.SidebarNavbarItem .SidebarNavbarItemLink .SidebarNavbarItemLinkHeading .SidebarNavbarItemLinkHeadingTitle:hover {
    color: var(--primaryColor);
}

 */


.SidebarNavbarItem .SidebarNavbarItemLink._Active .SidebarNavbarItemLinkHeading .SidebarNavbarItemLinkHeadingTitle {
    color: black;
}



.SidebarNavbarItem:hover > .SidebarNavbarItemLink .SidebarNavbarItemLinkHeadingTitle{
    color: #eee;
}

@media screen and (max-width: 960px) {
    .Sidebar {
        transform: translateX(-100%);
    }

    .Sidebar.MobileShow {
        transform: translateX(0%);
    }
}

</style>