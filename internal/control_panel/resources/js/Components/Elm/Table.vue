<script>
import {defineComponent} from 'vue'

export default defineComponent({
    name: "Table",
    props: {
        headers: {
            default: [],
            type: Array,
        },
        data: {
            default: [],
            type: Array,
        },
        btnActionShow: true
    },
    methods: {
        // Функция фильтрации items по ключам, присутствующим в headers
        filterItems(headers, items) {
            return items.map(item => {
                const newItem = {};
                Object.keys(item).forEach(key => {
                    if (headers.some(header => header.name === key)) {
                        newItem[key] = item[key];
                    }
                });
                return newItem;
            });
        },
        sortItemsByHeaders(itemColumnRows) {
            const filteredItems = this.filterItems(this.headers, Array(itemColumnRows));
            //
            return filteredItems.slice().sort((a, b) => {
                const keyA = Object.keys(a)[0];
                const keyB = Object.keys(b)[0];
                const indexA = this.headers.findIndex(header => header.name === keyA);
                const indexB = this.headers.findIndex(header => header.name === keyB);
                return indexA - indexB;
            });
        }
    },
    showItemService(item) {
        alert(JSON.stringify(item));
    }
})
</script>

<template>
    <div class="orderServiceMainList">
        <div class="orderServiceMainListWrapper">
            <div class="orderServiceMainListTable">
                <div class="tableHeaderColumn">
                    <div class="tableHeaderColumnHeaderWrapper"
                         style="">
                        <div class="tableHeaderRow" v-for="headItem in headers"
                             :class="`size_${headItem.size}`">
                            {{ headItem.title }}
                        </div>
                    </div>
                </div>
                <div class="tableOrderWrapper">
                    <div class="tableOrderColumn"
                         v-for="(itemColumnRows,columnKey) in data" :key="columnKey">
                        <!--                                                {{ sortItemsByHeaders(itemColumnRows) }}-->

                        <div class="orderColumnWrapper" style="display: flex;flex-flow: row;" v-if="itemColumnRows">
                            <template v-for="(item,keyRow) in sortItemsByHeaders(itemColumnRows)">
                                <slot name="itemRow" :item="item" :keyRow="keyRow">
                                    <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                                        {{ item }}
                                    </div>
                                </slot>
                            </template>

                            <template v-if="btnActionShow">
                                <slot name="itemRowBtn" v-if="itemColumnRows" :rowKey="columnKey">
                                    <div class="tableOrderRowAction size_1xs">
                                        <slot name="itemRowBtnItem" v-if="itemColumnRows" :rowItem="itemColumnRows" :rowKey="columnKey">
                                            <button @click="showItemService(itemColumnRows)"
                                                    style="">
                                                Управление
                                            </button>
                                        </slot>
                                    </div>
                                </slot>
                            </template>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<style scoped>
.tableHeaderColumn {
    background: #000;
    border-radius: 10px 10px 0px 0px;
    border: solid 1px #00FFB4FF;
    color: #00FFB4FF;
    opacity: 0.7;
    overflow: hidden;
}

.tableOrderColumn {
    border: solid 1px #222;
}

.tableOrderColumn:hover,
.tableOrderColumn:focus {
    border: solid 1px #217259;
    color: #00FFB4FF;
}

.tableHeaderColumnHeaderWrapper {
    display: flex;
    flex-flow: row;
    justify-content: left;
    align-items: center;
    background: #222;
}

.tableHeaderColumn .tableHeaderRow {
    border-right: solid 2px #333;
    padding: 12px 20px;
}

.tableHeaderColumn .tableHeaderRow:last-child {
    border-right: none;
    padding: 12px 20px;
}

.tableOrderColumn {
    /*padding: 14px 36px;*/
}

</style>


<style>
.size_1xs {
    min-width: 14vw;
}


.tableOrderColumn .tableOrderColumnRow {
    border-right: solid 2px #333;
    padding: 12px 20px;
    display: flex;
    align-items: center;
}

.tableOrderColumn .tableOrderRowAction {
    display: flex;
    align-items: center;
    padding: 12px 20px;
}

.tableOrderRowAction:last-child {
    min-width: unset;
    width: 100%;
}

.tableOrderRowAction button {
    padding: 4px 0;
    font-size: 14px;
    font-weight: 500;
    color: var(--primaryColor);
    background: none;
    border: none;
    border-radius: 6px;
    cursor: pointer;
}


</style>