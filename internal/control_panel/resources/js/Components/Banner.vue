<script>
import {defineComponent} from 'vue'

export default defineComponent({
    name: "Banner"
})
</script>

<template>

</template>

<style scoped>

/* service Banner */

.bannerServiceCard {
    /*margin-top: 16px;*/
    background: #EEFFFA;
    border-radius: 16px;
    width: 100%;
    box-sizing: border-box;
    display: flex;
    margin: 0 16px 24px;
}

.bannerServiceCardMid {
    min-height: 245px;
    flex: 1;
    width: 100%;
}

.bannerServiceCard .bannerServiceCardWrapper {
    padding: 24px 24px;
}

.bannerServiceCard .bannerServiceCardWrapperMid {
    padding: 48px 36px;
}

.bannerServiceCard .bannerServiceCardHeading .bannerServiceCardHeadingTitle {
    color: #02271D;
    font-family: 'Inter', sans-serif;
    font-size: 24px;
    font-style: normal;
    font-weight: 700;
    line-height: normal;
    margin-bottom: 24px;
}

.bannerServiceCard .bannerServiceCardHeading .bannerServiceCardHeadingDesc {
    color: #0B4F3C;
    font-family: 'Inter', sans-serif;
    font-size: 20px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
}

</style>