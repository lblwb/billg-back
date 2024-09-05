<script setup>
import {Head, router} from '@inertiajs/vue3';
import {onMounted} from "vue";
import Table from "@/Components/Elm/Table.vue";
import HeaderPanel from "@/Components/HeaderPanel.vue";

onMounted(async () => {

});

let PropsData = defineProps({
    services: Object | Array,
    // greeting: String,
    // users: Object | Array,
});

function getDateBdInfo(date) {
    let dateInfo = new Date(date);
    return dateInfo.toLocaleDateString();
}

let headersTable = [
    {title: 'Название', name: 'name', size: '1xs'},
    {title: 'Полное-название', name: 'full_name', size: '1xs'},
    {title: 'Тип-(-обр.-уcл.)', name: 'device_name', size: '1xs'},
    {title: 'Кол-во тарифов', name: 'tariffs', size: '1xs'},
    {title: 'Действие', name: 'btn', size: '1xs'},
];


</script>

<template>
    <Head title="Услуги"/>

  <!--        {{ PropsData.users }}-->

    <div class="services">

        <HeaderPanel titleHeader="Услуги" showBtnActionAdd="true" @clickAddContext="router.get('./add')"/>
        <!--            <div class="clientsItem" v-for="user in PropsData.users">-->
        <!--                {{ user }}-->
        <!--            </div>-->

        <Table :data="PropsData.services"
               :headers="headersTable" btnActionShow="true">

            <template v-slot:itemRow="{ item, keyRow }">
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.name }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.full_name }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.device_name }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    <div class="rowWrapper" style="display: flex; align-items: center;">
                        <div class="rowTitle" style="margin-right: 10px;">
                            {{ item && item.tariffs ? item.tariffs.length : "-" }} тр.
                        </div>
                        <div class="rowBtn">
                            <template v-if="item.tariffs !== null">
                                <div class="rowWrapper" style="display: flex; align-items: center;">
                                    <button class="btcAcceptPending"> ✏️ Добавить</button>
                                </div>
                            </template>
                        </div>
                    </div>
                </div>
            </template>
            <template v-slot:itemRowBtnItem="{rowItem}">
                <!--                {{ rowItem }}-->
                <button @click="router.get(`/panel/bill/services/${rowItem.slug}/info`)">Управление</button>
            </template>
        </Table>
    </div>
</template>

<style scoped>


.btcAcceptPending {
    display: flex;
    background: #222;
    color: #fff;
    border: none;
    padding: 6px 7px;
    font-size: 10px;
    border-radius: 10px;
    cursor: pointer;
}

.btcAcceptPending:hover,
.btcAcceptPending:focus {
    opacity: 0.7;
}
</style>
