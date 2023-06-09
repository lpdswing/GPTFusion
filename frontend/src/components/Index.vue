<script lang="ts" setup>
import {reactive, ref} from 'vue'
import { v4 as uuidv4 } from 'uuid'
import type { FormInstance, FormRules } from 'element-plus'
import {ReadMenu, EditMenu, WriteHome, ImportPlatfrom, ExportPlatfrom} from "../../wailsjs/go/main/App";
import {main} from "../../wailsjs/go/models"

const tableData = ref<main.PlatForm[]>([])

function init() {
    readMenu();
}

init();

function readMenu() {
    ReadMenu().then((data) => {
        tableData.value = data
    })
}

function editMenu(platforms: main.PlatForm[]) {
    EditMenu(platforms).then((data) => {
        readMenu();
    })
}


// 自定义模板列的方法
const edit = (row: main.PlatForm) => {
    // 编辑行
    dialogVisible.value = true
    dlgTitle.value = '编辑平台'
    newRowData.id = row.id
    newRowData.label = row.label
    newRowData.url = row.url
    newRowData.group = row.group
};
const removeId = ref('')
const del = (row: main.PlatForm) => {
    // 删除行
    removeId.value = row.id
    deleteDialogVisible.value = true

};

const confirm_del = () => {
    // 删除行
    tableData.value = tableData.value.filter((item) => item.id !== removeId.value)
    // 更新文件
    editMenu(tableData.value)
    deleteDialogVisible.value = false
};


// 弹出框可见状态和表单数据
const dlgTitle = ref('添加平台')
const dialogVisible = ref(false)
const deleteDialogVisible = ref(false)
const newRowData = reactive({
    id: '',
    label: '',
    url: '',
    priority: 0,
    separator: false,
    group: ""
})

const addRowFormRef = ref<FormInstance>()
// 表单验证规则
const newRowRules = reactive<FormRules>({
    label: [
        {required: true, message: '请输入平台名称', trigger: 'blur'},
        { min: 2, max: 10, message: 'Length should be 2 to 10', trigger: 'blur' }
    ],
    url: [
        {required: true, message: '请输入平台地址', trigger: 'blur'}
    ],
    group: [
        {required: false, message: '请输入分组名称', trigger: 'blur'}
    ]
})

// 添加行
const addRow = () => {
    // 表单验证
    addRowFormRef.value?.validate((valid) => {
        if (valid) {
            // 添加行
            console.log(newRowData)
            const idx = tableData.value.findIndex((item) => item.id === newRowData.id)
            if ( idx !== -1) {
                tableData.value[idx] = {...newRowData}
            } else {
                tableData.value.push({
                    id: uuidv4(),
                    label: newRowData.label,
                    url: newRowData.url,
                    priority: 0,
                    separator: false,
                    group: newRowData.group
                })
            }
            editMenu(tableData.value)
            // 关闭弹出框
            dialogVisible.value = false
            dlgTitle.value = '添加平台'
            // 重置表单
            addRowFormRef.value?.resetFields()
            newRowData.id = ''
        }
    })
}

const closeDig = () => {
    dialogVisible.value = false
    dlgTitle.value = '添加平台'
    addRowFormRef.value?.resetFields()
    newRowData.id = ''
    newRowData.label = ''
    newRowData.url = ''
    newRowData.group = ''
}

// localstorage
WriteHome(window.location.href).then((data) => {
    console.log(data)
})

const importPlatform = function () {
    ImportPlatfrom().then((data) => {
        console.log(data)
        readMenu();
    })
}

const exportPlatform = function () {
    ExportPlatfrom().then((data) => {
        console.log(data)
    })
}
</script>

<template>
    <div class="container">
        <el-text class="mx-1" type="warning">注意: Mac和Linux系统添加平台后需重启软件使菜单生效.</el-text>
        <div class="btn">
            <el-button style="margin: 10px 5px" @click="dialogVisible = true" type="primary" size="large">添加平台</el-button>
            <el-button style="margin: 10px 5px" @click="importPlatform" type="primary" size="large">导入平台</el-button>
            <el-button style="margin: 10px 5px" @click="exportPlatform" type="primary" size="large">导出平台</el-button>
        </div>
        <el-table :data="tableData" style="width: 100%; height: 100%">
            <template #empty>
                <div class="no-data">
                    <i class="iconfont icon-no-data"></i>
                    <p>暂无数据</p>
                </div>
            </template>
            <el-table-column prop="label" label="平台名称" width="150vw"/>
            <el-table-column prop="url" label="平台地址"/>
            <el-table-column prop="group" label="分组"/>
            <el-table-column label="操作" fixed="right" width="150vw">
                <template #default="{ row }">
                    <el-button type="primary" size="small" @click="edit(row)"
                    >编辑
                    </el-button
                    >
                    <el-button type="danger" size="small" @click="del(row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
    <el-dialog v-model="dialogVisible" :title="dlgTitle" draggable>
        <el-form ref="addRowFormRef" :model="newRowData" :rules="newRowRules">
            <!-- 表单项配置 -->
            <el-form-item label="平台名称" prop="label">
                <el-input v-model="newRowData.label" placeholder="请输入平台名称: 比如(百度)"/>
            </el-form-item>
            <el-form-item label="平台链接" prop="url">
                <el-input v-model="newRowData.url" placeholder="请输入平台链接: 比如(https://www.baidu.com)"/>
            </el-form-item>
            <el-form-item label="平台分组" prop="group">
                <el-input v-model="newRowData.group" placeholder="请输入分组"/>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="closeDig">取消</el-button>
                <el-button type="primary" @click="addRow">确定</el-button>
            </span>
    </el-dialog>
    <el-dialog title="确认删除" v-model="deleteDialogVisible" width="50%" draggable="true">
        <span style="color: black;">删除后将不可恢复，确定要删除吗？</span>
        <div style="margin-top: 36px;" slot="footer" class="dialog-footer">
            <el-button @click="deleteDialogVisible = false">取消</el-button>
            <el-button type="danger" @click="confirm_del">确定</el-button>
        </div>
    </el-dialog>
</template>

<style scoped>
.container {
    /*height: 100%;*/
    width: 100%;
    display: flex;
    flex-direction: column;
}

.container .btn {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    background-color: #fff;
}
</style>
