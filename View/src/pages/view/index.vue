<template>
    <div>
        <p>{{ res }}</p>
        <button @click="getString">文字列を取得</button>
        <button @click="getProfile">プロフィールを取得</button>
    </div>
</template>
<script>

export default {

    data: function() {
        return {
            res: '選択されていません' 
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            console.log('load')
        })
    },
    methods: {
        async getString() {
            this.res = await this.$axios.$get('/resString')
        },
        async getProfile() {
            let res = await this.$axios.$get('/getProfile',{
                params:{}
            })
            .then(response => {
                this.res = response.profile
            })
            .catch(err => {
                console.log(err)
            })
            .finally(() => console.log('finally'))
        }
    }
}

// axios.get(url).then(res => {
//   const items = res.data;
//   for (const item of items) {
//     console.log(`${item.user.id}: \t${item.title}`);
//   }
// }).catch(error => {
//   const {
//     status,
//     statusText
//   } = error.response;
//   console.log(`Error! HTTP Status: ${status} ${statusText}`);
// });

</script>