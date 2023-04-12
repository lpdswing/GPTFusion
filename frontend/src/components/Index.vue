<script lang="ts" setup>
import {reactive, ref} from 'vue'
import { v4 as uuidv4 } from 'uuid'
import type { FormInstance, FormRules } from 'element-plus'
import { ReadMenu, EditMenu } from "../../wailsjs/go/main/App";
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
};

const del = (row: main.PlatForm) => {
    // 删除行
    tableData.value = tableData.value.filter((item) => item.id !== row.id)
    // 更新文件
    editMenu(tableData.value)
};


// 弹出框可见状态和表单数据
const dlgTitle = ref('添加平台')
const dialogVisible = ref(false)
const newRowData = reactive({
    id: '',
    label: '',
    url: ''
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
                    url: newRowData.url
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
}
</script>

<template>
    <div class="container">
        <div class="btn">
            <el-button style="margin: 10px 5px" @click="dialogVisible = true" type="primary" size="large">添加平台</el-button>
        </div>
        <el-table :data="tableData" style="width: 100%; height: 100%">
            <template #empty>
                <div class="no-data">
                    <i class="iconfont icon-no-data"></i>
                    <p>暂无数据</p>
                </div>
            </template>
            <el-table-column prop="label" label="平台名称" width="150"/>
            <el-table-column prop="url" label="平台地址"/>
            <el-table-column label="操作" fixed="right" width="150">
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
                <el-input v-model="newRowData.label" />
            </el-form-item>
            <el-form-item label="平台链接" prop="url">
                <el-input v-model="newRowData.url" />
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="closeDig">取消</el-button>
                <el-button type="primary" @click="addRow">确定</el-button>
            </span>
    </el-dialog>
</template>

<style scoped>
.container {
    height: 100%;
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
