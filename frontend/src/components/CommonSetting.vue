<template>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <span>模式选择</span>
            </div>
        </template>
        <div>
            <el-radio-group v-model="mode_radio" @change="mode_change">
                <el-radio label="1" size="large" border>窗口模式</el-radio>
                <el-radio label="2" size="large" border>侧边栏模式</el-radio>
            </el-radio-group>
        </div>
    </el-card>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <span>窗口选项</span>
            </div>
        </template>
        <el-switch
                v-model="alwaysTop"
                class="mb-2"
                size="large"
                @change="onAlwaysTopChange"
                active-text="窗口始终置顶"
                inactive-text="取消窗口置顶"
        />
        <br>
        <el-switch
                v-model="hideWindowOnClose"
                class="mb-2"
                size="large"
                @change="onHideWindowOnCloseChange"
                active-text="点击关闭按钮隐藏窗口"
                inactive-text="点击关闭按钮关闭窗口"
        />
        <br>
        <el-text class="mx-1" type="warning">注意: 此设置重启软件生效.</el-text>
    </el-card>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {main} from "../../wailsjs/go/models"
import {ReadSetting, WriteSetting} from "../../wailsjs/go/main/App";

let setting = ref<main.Setting>({
    mode: "1",
    always_on_top: true,
    hide_window_on_close: true,
})

const mode_radio = ref("1")
// 始终置顶
const alwaysTop = ref(false)
// 关闭隐藏
const hideWindowOnClose = ref(true)

ReadSetting().then((data) => {
    setting.value = data

    mode_radio.value = setting.value.mode
    alwaysTop.value = setting.value.always_on_top
    hideWindowOnClose.value = setting.value.hide_window_on_close
})

const mode_change = (val: string) => {
    setting.value.mode = val
    WriteSetting(setting.value)
}


const onAlwaysTopChange = (val: boolean) => {
    setting.value.always_on_top = val
    WriteSetting(setting.value)
}

const onHideWindowOnCloseChange = (val: boolean) => {
    setting.value.hide_window_on_close = val
    WriteSetting(setting.value)
}

</script>

<style scoped>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 15px;
}

.box-card {
    width: 100%;
    margin-top: 20px;
}
</style>