<script lang="ts" setup>
import Index from './components/Index.vue'
import { ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'
import CommonSetting from "./components/CommonSetting.vue";
import {GetVersion, WriteHome} from "../wailsjs/go/main/App";


const activeName = ref('first')
const appVersion = ref('unknown')
GetVersion().then((data) => {
    appVersion.value = data
})
const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}
WriteHome(window.location.href).then((data) => {
    console.log(data)
})
</script>

<template>
    <div class="mycontent">
        <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="通用设置" name="first"><CommonSetting /></el-tab-pane>
            <el-tab-pane label="平台管理" name="second"><Index /></el-tab-pane>
            <footer class="footer">© Github@lpdswing Version {{ appVersion }} Powered By wails </footer>
        </el-tabs>
    </div>
</template>

<style>
.mycontent {
    height: 100%;
    background-color: #fff;
}
.demo-tabs {
    padding: 20px 20px;
    background-color: #fff;
}
.demo-tabs > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
    background-color: #fff;
}
.footer {
    height: 25px;
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    font-size: small;
    color: rgba(0, 0, 0, 0.3);
    text-align: center;
    z-index: 999;
}
</style>
