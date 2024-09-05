<script setup>
import {Head, router} from '@inertiajs/vue3';
import {onMounted} from "vue";
import Table from "@/Components/Elm/Table.vue";
import HeaderPanel from "@/Components/HeaderPanel.vue";

// const hiddenRow = ['order_info', 'service_about', 'service_status', 'service_tariff', 'service_info', 'service_price', 'vw', 'created_at', 'updated_at'];
const showRow = ['service_info', 'service_tariff', 'created_at', 'service_status'];

onMounted(async () => {

});


function filteredItem(item, field) {
    return hiddenRow[field];
}

function getStatus(statusName) {
    switch (statusName) {
        case "accepted":
            return "Одобрен"
        case "pending":
            return "Ожидает подтверждения"
        case "pay_pending":
            return "Ожидает оплаты"
        default:
            return "Не задан!"
    }
}

function getDateInfo(date) {
    let dateInfo = new Date(date);
    return dateInfo.toLocaleString();
}


let PropsData = defineProps({
    userSvcOrders: Object | Array,
    // greeting: String,
    // users: Object | Array,
});

let headersTable = [
    {title: 'Услуга', name: 'service_info', size: '1xs'},
    {title: 'Ресурс', name: 'resource', size: '1xs'},
    {title: 'Сумма заказа', name: 'order_info', size: '1xs'},
    {title: 'Статус', name: 'service_status', size: '1xs'},
    // {title: 'Клиент', name: 'order_info', size: '1xs'},
    {title: 'Дата создания', name: 'created_at', size: '1xs'},
    {title: 'Действие', name: 'btn', size: '1xs'},
];

</script>

<template>
    <Head title="Заказы клиентов"/>
    <div class="orders">

        <HeaderPanel titleHeader="Заказы" showBtnActionAdd="true" @clickAddContext="router.get('./add')"/>

        <!--        {{ PropsData.userSvcOrders }}-->

        <Table :data="PropsData.userSvcOrders"
               :headers="headersTable">

            <template v-slot:itemRow="{ item, keyRow }">

                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.service_info.name }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.resource }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.order_info.total_amount }} руб.
                </div>

                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    <div class="rowWrapper" style="display: flex; align-items: center;">
                        <div class="rowTitle" style="margin-right: 6px;">
                            {{ getStatus(item.service_status) }}
                        </div>
                        <div class="rowBtn">
                            <template v-if="item.service_status === 'pending'">
                                <div class="rowWrapper" style="display: flex; align-items: center;">
                                    <button class="btcAcceptPending" style="margin-right: 3px;"> 🧾 Возврат</button>
                                    <button class="btcAcceptPending"> ✅ Принять</button>
                                </div>
                            </template>
                        </div>
                    </div>
                </div>

                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow" style="max-width: 80px">
                    {{ getDateInfo(item.created_at) }}
                </div>
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
