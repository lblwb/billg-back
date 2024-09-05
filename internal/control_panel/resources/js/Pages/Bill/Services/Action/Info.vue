<script>
import {defineComponent, toRaw} from 'vue'
import {Head} from "@inertiajs/vue3";

export default defineComponent({
    name: "BillSvcActionInfo",
    props: {
        serviceData: Object | Array,
        serviceId: String,
    },
    methods: {
        async findIndexByOptionName(jsonData, optionName) {
            for (const index in jsonData) {
                if (jsonData[index].option_name === optionName) {
                    return index;
                }
            }
            return -1; // If the option_name is not found
        },
        async tariffDelParam(tariffItem, paramItem) {
            const confirmed = confirm("Вы уверены что хотите удалить данный параметр!?");
            if (confirmed) {
                // let tarParams = toRaw(tariffItem).params;
                let tarParams = {...tariffItem.params}; // Make a shallow copy of params
                console.log(tarParams);

                // console.log(tariffItem.params, paramItem.option_name);
                const currentIndex = await this.findIndexByOptionName(toRaw(tarParams), paramItem.option_name);
                delete tarParams[currentIndex];
                tariffItem.params = tarParams;

                if (currentIndex !== -1) {
                    delete tarParams[currentIndex];
                    tariffItem.params = tarParams;
                    alert("удаление " + currentIndex);
                } else {
                    alert("Параметр не найден.");
                }

                // alert("удаление" + currentIndex);
            } else {

            }
        },
        async tariffAddParam(tariffItem) {
            alert("Добавление параметра!" + JSON.stringify(tariffItem));
        },
        tariffEditParam(tariffItem, paramItem) {
            alert(JSON.stringify(paramItem));
        }
    },
    components: {
        Head
    }
})
</script>

<template>

    <Head :title="`Информация об услуге - [[${serviceData.name}]]`"/>

    <div class="serviceInfo" v-if="serviceData">
        <div class="serviceInfoHeader" style="margin-bottom: 36px;">
            <div class="serviceInfoHeaderTitle" style="font-size: 20px; font-weight: 600; text-transform: uppercase;">
                Услуга: [[{{ serviceData.full_name }}]] [[{{ serviceId }}]]
            </div>
        </div>
        <div class="serviceInfoBody">
            <div class="serviceInfoBodyInfoList" style="display: flex">
                <div class="serviceInfoBodyInfoListItem" style="margin-bottom: 24px">

                    <!--                    <template v-if="key !== 'tariffs'" style="flex: auto">-->
                    <div class="infoListItemInfo">
                        <div class="infoListItem">
                            <div class="infoListItemWrapper">
                                <div class="infoListItemLabel">
                                    Слаг:
                                </div>
                                <div class="infoListItemValue">
                                    {{ serviceData.slug }}
                                </div>
                            </div>
                        </div>

                        <div class="infoListItem">
                            <div class="infoListItemWrapper">
                                <div class="infoListItemLabel">
                                    Полное-название:
                                </div>
                                <div class="infoListItemValue">
                                    {{ serviceData.full_name }} / {{ serviceData.full_name_en }}
                                </div>
                            </div>
                        </div>
                        <div class="infoListItem">
                            <div class="infoListItemWrapper">
                                <div class="infoListItemLabel">
                                    Название-оборудования:
                                </div>
                                <div class="infoListItemValue">
                                    {{ serviceData.device_name }}
                                </div>
                            </div>
                        </div>

                        <div class="infoListItem">
                            <div class="infoListItemWrapper">
                                <div class="infoListItemLabel">
                                    Слаг-оборудования:
                                </div>
                                <div class="infoListItemValue">
                                    {{ serviceData.device_slug }}
                                </div>
                            </div>
                        </div>
                        <div class="infoListItem">
                            <div class="infoListItemWrapper">
                                <div class="infoListItemLabel">
                                    DeviceSlug:
                                </div>
                                <div class="infoListItemValue">
                                    {{ serviceData.device_slug }}
                                </div>
                            </div>
                        </div>
                        <div class="infoListItem">
                            <div class="infoListItemWrapper" style="flex-flow: column;">
                                <div class="infoListItemLabel" style="margin-bottom: 14px">
                                    Текст баннера:
                                </div>
                                <div class="infoListItemValue">
                                    <div class="infoListItemValueWrapper">
                                        {{ serviceData.banner_desc }}
                                        <hr>
                                        {{ serviceData.banner_desc_en }}
                                    </div>
                                    <!--                                    {{ serviceData.banner_desc }} / {{ serviceData.banner_desc_en }}-->
                                </div>
                            </div>
                        </div>
                    </div>
                    <!--                    </template>-->


                    <!--                    <hr>-->
                    <!--                    {{ serviceData }}-->

                    <div class="serviceInfoBodyTariffs" v-if="serviceData.tariffs">
                        <div class="serviceInfoBodyTariffsHeader" style="margin-bottom: 14px">
                            <div class="serviceInfoBodyTariffsHeaderTitle"
                                 style="font-size: 14px; font-weight: 600; text-transform: uppercase;">
                                Тарифные планы:
                                <button class="miniBtnTariff">🎟️ Добавить тариф</button>
                            </div>
                        </div>
                        <div class="serviceInfoBodyTariffsBody">
                            <div class="infoBodyTariffsBody"
                                 style="">

                                <div class="infoBodyTariffsBodyWrapper"
                                     style="display: flex; flex-flow: row wrap; min-width: 48vw;">
                                    <div class="infoBodyTariffsBodyInfoListItem"
                                         v-for="(itemTariff) in serviceData.tariffs"
                                         style="width: 100%;padding: 10px; border: solid 1px #222; border-radius: 14px; margin-bottom: 12px;max-width: 20.5vw; margin-right: 14px;">
                                        <div class="bodyInfoListItem" style="margin-bottom: 8px; font-weight: 600">
                                            Название: {{ itemTariff.full_name }} / {{ itemTariff.full_name_en }} | Слаг:
                                            {{ itemTariff.slug }}
                                        </div>
                                        <div class="bodyInfoListItem" style="margin-bottom: 8px">
                                            <span>Описание тарифа: </span><span> {{
                                            itemTariff.desc_alert
                                            }} / {{ itemTariff.desc_alert_en }}</span>
                                        </div>

                                        <hr>

                                        <div class="bodyInfoListItem">
                                            <div class="bodyInfoListItemHeader" style="margin-bottom: 8px">Параметры
                                                тарифа:
                                                <button class="miniBtnTariff" @click="tariffAddParam(itemTariff)">⚙️
                                                    Добавь
                                                    параметр
                                                </button>
                                            </div>
                                            <div class="bodyInfoListItemParams">
                                                <div class="bodyInfoListItemParamsItem"
                                                     v-for="param in itemTariff.params"
                                                     style="margin-bottom: 4px; padding: 8px; border-bottom: solid 1px #333; background: #222; border-radius: 10px;">
                                                    <!--  -->
                                                    <div class="itemParamsItemWrapper"
                                                         @click="tariffEditParam(itemTariff,param)"
                                                         style="cursor:pointer;">
                                                        <div class="bodyInfoListItemParamsItemHead">
                                                            Поле: {{ param.option_name }} / {{ param.params }}
                                                        </div>
                                                        <div class="bodyInfoListItemParamsItemHead">
                                                            Цена: {{ param.price_unit }}
                                                        </div>

                                                        <div class="bodyInfoListItemParamsItemHead">
                                                            Тип:
                                                            <span v-if="param && param.type_param === 'input'">Поле</span>
                                                            <span v-else-if="param && param.type_param === 'range'">Слайдер</span>
                                                        </div>
                                                    </div>

                                                    <button class="miniBtnTariff"
                                                            @click="tariffDelParam(itemTariff,param)">🔴 Удалить
                                                    </button>

                                                </div>
                                            </div>

                                        </div>
                                    </div>
                                </div>

                            </div>
                        </div>
                    </div>

                    <!--                        {{ key }}-->

                </div>
            </div>
        </div>
    </div>
  <!--  {{ serviceData }}-->
  <!--  {{ serviceId }}-->
</template>

<style scoped>

.infoListItemWrapper {
    display: flex;
}

.infoListItemInfo {
    display: flex;
    flex-flow: row wrap;
    max-width: 28vw;
    margin-bottom: 24px;
    align-items: flex-start;
}

.infoListItem {
    display: inline-flex;
    margin-bottom: 14px;
    border: solid 1px #333;
    border-radius: 10px;
    margin-right: 18px;
    padding: 8px;
}

.serviceInfoBodyInfoListItem .infoListItemLabel {
    font-size: 14px;
    font-weight: 500;
    margin-right: 6px;
}

.serviceInfoBodyInfoListItem .infoListItemValue {
    font-size: 14px;
    font-weight: 600;
}

.miniBtnTariff {
    font-size: 11px;
    background: #222;
    color: #fff;
    border-radius: 6px;
    padding: 3px 6px;
    border: solid 1px #666;
    margin: 4px 0;
}
</style>