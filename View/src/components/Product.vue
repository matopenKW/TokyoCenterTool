<template>
    <div class="inner">
        <h2 class="title">PRODUCT</h2>
        <div v-for="product in products" :key="product.Title">
            <div class="box11">
                <p class="img"></p>
                <p class="p-header">{{ product.Title }}</p>
                <p>{{ product.Detail}}</p>
                <p><nuxt-link to="/portfolio">nuxt-linkのリンク</nuxt-link></p>
            </div>
        </div>
    </div>
</template>

<script>

export default {

    data: function(){
        return {
            products: {}
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            this.$axios.$get('/getProduct',{
                params:{}
            })
            .then(response => {
                this.products = response.product
            })
            .catch(err => {
                console.log(err)
            })
            .finally(() => console.log('finally'))
        })
    },
    computed: {
        hogeURL: function () {
            return product.ProductUrl
        }
    }

}

</script>