<script>
import {defineComponent} from 'vue';

export default defineComponent({
    name: "AdBillNavbar",
    data() {
        return {
            userName: "admin: Username"
        }
    },
    methods: {},
    computed: {
        getTitleName() {
            return window.document.title;
        }
    }
})
</script>

<template>
    <div class="BodyNavbar">
        <div class="BodyNavbarWrapper">
            <div class="BodyNavbarRow">
                <svg width="64" height="2" viewBox="0 0 64 2" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M0 1H64" stroke="#AAF9E1"/>
                </svg>
                {{ getTitleName }}
                <!-- TODO: HTML / NAVBAR - Тут скоро что-то будет! -->
            </div>
            <div class="BodyNavbarRow">
                <div class="BodyNavbarUserAccount">
                    <div class="BodyNavbarUserAccountWrapper">
                        <div class="BodyNavbarUserProfile" style="margin-right: 16px">
                            <div class="BodyNavbarUserProfileBtn">
                                <!--                                 @click="$router.push({name: 'user-account'})"-->
                                <div class="BodyNavbarUserProfileHeading">
                                    <div class="BodyNavbarUserProfileHeadingTitle" style="">
                                        {{ userName }}
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="BodyNavbarUserLogout">
                            <div class="BodyNavbarUserLogoutBtn">
                                <button style="color: #fff; background: transparent; border: solid 2px #222; padding: 8px 10px; border-radius: 12px; height: 100%; cursor: pointer">
                                    <span class="BodyNavbarUserLogoutBtnIcon">
                                        🫡 Выход
<!--                                        <img src="~/assets/images/icon/logout.svg">-->
                                    </span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>


<style scoped>

/*Navbar*/

.BodyNavbar {

}

.sticky {
    position: fixed;
    top: 0;
    background-color: #fff;
    box-shadow: 0px 3px 0px rgb(0 0 0 / 5%);
    z-index: 1000;
    margin: 0 auto;
    right: 0;
    left: 320px;
}

@media screen and (max-width: 960px) {
    .sticky {
        left: 0;
    }
}

.BodyNavbarWrapper {
    padding: 16px 24px;
    display: flex;
    justify-content: space-between;
}

@media screen and (max-width: 960px) {
    .BodyNavbar {
        width: 100%;
    }

    .BodyNavbarWrapper {
        padding: 14px 24px;
        width: 100%;
        box-sizing: border-box;
    }

    .BodyNavbarWrapper .BodyNavbarRow:first-child {
        display: none;
    }

    .BodyNavbarRow {
        flex-flow: row wrap;
        justify-content: left !important;
    }
}

.BodyNavbarRow {
    display: flex;
    align-items: center;
    flex: auto;
}

.BodyNavbarRow:last-child {
    display: flex;
    justify-content: flex-end;
}

.BodyNavbarUserAccount {

}

.BodyNavbarUserAccountWrapper {
    display: flex;
    align-items: center;
}

@media screen and (max-width: 960px) {
    .BodyNavbarUserAccountWrapper {
        flex-flow: row wrap;
    }
}

.BodyNavbarUserBalance {
    margin-right: 24px;
}


[data-theme="dark"] .BodyNavbarUserProfileHeadingTitle {

}

.BodyNavbarUserProfileBtn {
    border-radius: 20px;
    /* background: #F9F8F8;*/
    padding: 10px 26px;
    cursor: pointer;
    border: solid 2px #222;
}

[data-theme="dark"] .BodyNavbarUserProfileBtn {
    background: var(--bg-color);
}

.BodyNavbarUserProfileHeadingTitle {
    font-family: 'Inter', sans-serif;
    font-size: 12px;
    font-style: normal;
    font-weight: 500;
    line-height: normal;
}
</style>