<template>
    <div class="inner">
        <h2 class="title">PROFILE</h2>
        <div class="text">
            <div class="column-image">
                <img src="https://pbs.twimg.com/profile_images/1164314556631592960/xXEg31M0_400x400.jpg">
            </div>
            <div class="column-text">
                <p>{{ profile.Detail }}</p>
            </div>
        </div>
    </div>
</template>

<script>

export default {

    data: function(){
        return {
            profile: {}
        }
    },
    mounted: function () {
        this.$nextTick(() => {
            this.$axios.$get('/getProfile',{
                params:{}
            })
            .then(response => {
                this.profile = response.profile
            })
            .catch(err => {
                console.log(err)
            })
            .finally(() => console.log('finally'))
        })
    }
}

</script>