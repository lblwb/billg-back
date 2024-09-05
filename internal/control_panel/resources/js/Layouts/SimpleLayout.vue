<template>
    <div class="container">
        <AdBillSidebar/>
        <!--        <AdBillNavbar/>-->
        <div class="Body">
            <div class="bodyHeader">
                <AdBillNavbar/>
            </div>
            <div class="BodyContext">
                <slot></slot>
            </div>
        </div>
    </div>
</template>
<script setup>

import AdBillNavbar from "@/Components/AdBillNavbar.vue";
import AdBillSidebar from "@/Components/AdBillSidebar.vue";
</script>

<style>
@import url("../../css/billing/main.css");
</style>

<style scoped>

body {
    height: 100% !important;
}

.containerAndSidebar {
    width: 100%;
    display: flex;
    justify-content: space-between;
}

.containerWrapper {
    width: 100%;
    display: flex;
}


.Body {
    position: absolute;
    left: 320px;
    right: 0;
    top: 0;
    width: auto;
    bottom: auto;
    box-sizing: border-box;
}

@media screen and (max-width: 960px) {
    .Body {
        position: relative;
        right: 0;
        left: 0;
        margin: 0;
        padding: 0;
    }
}


.BodyContext {
    box-sizing: border-box;
    position: relative;
}


.BodyContext {
    padding: 0 24px;
    height: 100%;
}
</style>