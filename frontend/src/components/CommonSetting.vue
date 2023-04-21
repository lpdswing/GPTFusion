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
        <el-row :gutter="18" justify="start">
            <el-col :span="18">
                <el-switch
                        v-model="alwaysTop"
                        size="large"
                        @change="onAlwaysTopChange"
                        active-text="保持窗口置顶"
                />
            </el-col>
        </el-row>
        <el-row :gutter="18" justify="start">
            <el-col :span="18">
                <el-switch
                    v-model="rememberLastPage"
                    size="large"
                    @change="onRememberLastPageChange"
                    active-text="记住上次页面"
                />
            </el-col>
        </el-row>
        <el-row :gutter="18" justify="center">
            <el-col :span="18">
                <el-switch
                        v-model="hideWindowOnClose"
                        size="large"
                        @change="onHideWindowOnCloseChange"
                        active-text="关闭隐藏窗口"
                />
            </el-col>
            <el-col :span="6">
                <el-tooltip content="此项需重启生效" placement="top">
                    <el-button>提示</el-button>
                </el-tooltip>
            </el-col>
        </el-row>

    </el-card>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {main} from "../../wailsjs/go/models"
import {ReadSetting, WriteSetting} from "../../wailsjs/go/main/App";

let setting = ref<main.Setting>({
    mode: "1",
    always_on_top: false,
    hide_window_on_close: false,
    remember_last_page: true,
    last_page: "/",
})

const mode_radio = ref("1")
// 始终置顶
const alwaysTop = ref(false)
// 关闭隐藏
const hideWindowOnClose = ref(false)
const rememberLastPage = ref(false)

ReadSetting().then((data) => {
    setting.value = data

    mode_radio.value = setting.value.mode
    alwaysTop.value = setting.value.always_on_top
    hideWindowOnClose.value = setting.value.hide_window_on_close
    rememberLastPage.value = setting.value.remember_last_page
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

const onRememberLastPageChange = (val: boolean) => {
    setting.value.remember_last_page = val
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