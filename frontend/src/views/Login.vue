<template>
    <div>
      <el-row style="margin-top: 10px">
                <el-col>
                    <el-card shadow="always" style="max-width: 48rem;margin:0 auto">
                        <div slot="header">
                            Shevon 订阅统一控制面板 登录

                        </div>   <el-container>
                                <el-form :model="form" label-width="80px" label-position="left" style="width: 100%">
                                   <el-form-item label="用户名:">
                                        <el-input v-model="form.user" type="username" />
                                    </el-form-item>
                                    <el-form-item label="密码:">
                                            <el-input v-model="form.password" type="password"/>
                                        </el-form-item>
                                        <el-form-item label-width="0px" style="text-align: right">

                                    <el-button plain style="width: 30%" type="success" @click="login">登录</el-button>
                                </el-form-item>
                                    </el-form>
                                    </el-container>
                        </el-card>  
                       </el-col> 
                     
                       </el-row>


    </div>
</template>

<script>
export default {
    data() {
        return {
            form: {
                user: '',
                password: '',
            },
        };
    },
    methods: {
        login() {
            this.$axios
                .post('/api/login', {
                    user: this.form.user,
                    password: this.form.password,
                })
                .then(res => {
                    if (res.data.code == 1 && res.data.token != "") {
                        this.setLocalStorageItem("token", res.data.token);
                        this.$message.success("登录成功");
                        this.$router.push('/controller')
                    } else {
                        this.$message.error("账号或密码错误");
                    }
                })
                .catch(() => {
                    this.$message.error("未知错误");
                });
        },
        getLocalStorageItem(itemKey) {
            const now = +new Date()
            let ls = localStorage.getItem(itemKey)

            let itemValue = ''
            if (ls !== null) {
                let data = JSON.parse(ls)
                if (data.expire > now) {
                    itemValue = data.value
                } else {
                    localStorage.removeItem(itemKey)
                }
            }

            return itemValue
        },
        setLocalStorageItem(itemKey, itemValue) {
            const ttl = process.env.VUE_APP_CACHE_TTL
            const now = +new Date()

            let data = {
                setTime: now,
                ttl: parseInt(ttl),
                expire: now + ttl * 1000,
                value: itemValue
            }
            localStorage.setItem(itemKey, JSON.stringify(data))
        }
    },
};
</script>
