<script setup>
import {Head, router} from '@inertiajs/vue3';
import {onMounted} from "vue";
import Table from "@/Components/Elm/Table.vue";
import HeaderPanel from "@/Components/HeaderPanel.vue";

onMounted(async () => {

});

let PropsData = defineProps({
    users: Object | Array,
    // greeting: String,
    // users: Object | Array,
});

function getDateBdInfo(date) {
    let dateInfo = new Date(date);
    return dateInfo.toLocaleDateString();
}

let headersTable = [
    {title: 'Ник', name: 'username', size: '1xs'},
    {title: 'Баланс', name: 'balance', size: '1xs'},
    {title: 'TG-ID', name: 't_id', size: '1xs'},
    {title: 'Кол-во заказов', name: 'orders', size: '1xs'},
    {title: 'Дата рождения', name: 'birthday', size: '1xs'},
    {title: 'Действие', name: 'btn', size: '1xs'},
    // {title: 'Ресурс', name: 'resource', size: '1xs'},
    // {title: 'Статус', name: 'service_status', size: '1xs'},
    // {title: 'Дата создания', name: 'created_at', size: '1xs'},
    // {title: 'Действие', name: 'btn', size: '1xs'},
];


</script>

<template>
    <Head title="Клиенты"/>

  <!--        {{ PropsData.users }}-->

    <div class="clients">

        <HeaderPanel titleHeader="Клиенты" showBtnActionAdd="true" @clickAddContext="router.get('./add')"/>

        <Table :data="PropsData.users"
               :headers="headersTable">

            <template v-slot:itemRow="{ item, keyRow }">
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.username }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.balance ? item.balance.amount + " руб." : "-" }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.t_id ? item.t_id : "-" }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item && item.orders ? item.orders.length + " зак." : "-" }}
                </div>
                <div class="tableOrderColumnRow size_1xs" :data-key="keyRow">
                    {{ item.birthday ? getDateBdInfo(item.birthday) : "-" }}
                </div>
            </template>
        </Table>
    </div>
</template>

<style scoped>
</style>
