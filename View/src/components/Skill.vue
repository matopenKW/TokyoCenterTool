<template>
    <div class="inner">
        <h2 class="title">SKILL</h2>
        <div v-for="skill in skills" :key="skill.id">
            <div class="text box-row">
                <div class="box back-end">
                    <div class="box-title">{{ skill.name }}</div>
                    <p>Java ４年<br>
                        [フレームワーク]<br>
                        SpringBoot、Spring、Struts2、Seasar2、Play framework</p>
                    <p>Swift 0.5年<br>
                        [開発内容]<br>
                        iosアプリ開発。チャットアプリケーション
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {

    data: function(){
        return {
            skills: []
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            this.$axios.$get('/getSkill',{
                params:{}
            })
            .then(response => {
                
                //this.skill = response.skill
                response.skill.forEach(element => {
                    var skill = {
                        id : element.CategoryId,
                        name : element.CategoryName
                    }
                    this.skills.push(skill);
                });
            })
            .catch(err => {
                console.log(err)
            })
            .finally(() => console.log('finally'))
        })
    }
}

</script>