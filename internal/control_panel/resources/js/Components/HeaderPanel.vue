<script>
import {defineComponent} from 'vue'

export default defineComponent({
    name: "HeaderPanel",
    props: {
        titleHeader: {
            default: ""
        },
        showBtnActionAdd: {
            default: true
        },
        clickAddContext: {}
    },
    methods: {
        addNewContext() {
            this.$emit("clickAddContext");
        }
    }
})
</script>

<template>
    <div class="HeaderPanelInfoHeader" style="margin-bottom: 36px;">
        <div class="HeaderPanelInfoHeaderWrapper" style="display: flex; flex-flow: row wrap; align-items: center;">
            <div class="HeaderPanelInfoHeaderTitle"
                 style="font-size: 20px; font-weight: 600; text-transform: uppercase; flex: auto;">
                {{ titleHeader }}
            </div>
            <div class="HeaderPanelInfoHeaderAction" v-if="showBtnActionAdd">
                <div class="infoHeaderActionBtn">
                    <button style="border: solid 1px #222; background: none; color: #fff; padding: 6px 10px; border-radius: 10px"
                            @click="addNewContext">🔥 Добавить
                    </button>
                </div>
            </div>
        </div>

    </div>
</template>

<style scoped>

</style>