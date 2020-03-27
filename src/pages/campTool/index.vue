<template>
    <div id="content">
        <div class="nav"></div>
        <div class="container-fluid">
            <div v-for="file in files" :key="file.id" class="row">
                <div class="text-center col-3">
                    <p> Id: {{ file.id}} Name: {{file.Name}}</p>
                </div>
            </div>
            <div class="row">
                <div class="col-12">
                    <button @click="getString">選択したファイルを名寄せする</button>
                </div>
            </div>
        </div>
        <div class="container-fluid">
            <div class="row">
                <div class="col-12">
                    <p>名寄せファイル</p>
                </div>
            </div>
            <div class="row">
                <div class="col-12">
                    <button @click="getString">ダウンロード</button>
                </div>
            </div>
        </div>
        <div>
            <div>
                <p>{{ res }}</p>
                <a
                href="https://github.com/nuxt/nuxt.js"
                target="_blank"
                class="button--grey"
                >link</a>
            </div>
            <div>
                <button @click="getString">文字列を取得</button>
                <button @click="getProfile">プロフィールを取得</button>
            </div>
        </div>
    </div>
</template>
<script>

export default {

    data: function() {
        return {
            files: [
                {
                    id: 1,
                    Name: "雪まつり名簿",
                },
                {
                    id: 1,
                    Name: "申し込み",
                },
                {
                    id: 1,
                    Name: "申し込み",
                },
                {
                    id: 1,
                    Name: "申し込み",
                },
                {
                    id: 1,
                    Name: "申し込み",
                }
            ]
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            console.log('load')
        })
    },
    methods: {
        async clickFile(){
            alert('click File')

        }


        // async getString() {
        //     this.res = await this.$axios.$get('/resString')
        // },
        // async getProfile() {
        //     let res = await this.$axios.$get('/getProfile',{
        //         params:{}
        //     })
        //     .then(response => {
        //         this.res = response.profile
        //     })
        //     .catch(err => {
        //         console.log(err)
        //     })
        //     .finally(() => console.log('finally'))
        // }
    }
}

</script>