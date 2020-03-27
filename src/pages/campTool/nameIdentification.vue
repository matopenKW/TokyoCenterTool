<template>
    <div id="content" class="p-3">
        <div class="nav"></div>
        <div class="container-fluid">
            <div class="row">
                <div class="col-3">
                    <h3>フォルダ１</h3>
                    <div class="row folder-list">
                        <div v-for="file in files1" :key="file.id" @click="clickFile" class="col-12">{{file.Name}}</div>
                    </div>
                </div>
                <div class="col-3">
                    <h3>フォルダ２</h3>
                    <div class="row folder-list">
                        <div v-for="file in files2" :key="file.id" @click="clickFile" class="col-12">{{file.Name}}</div>
                    </div>
                </div>
                <div class="col-2">
                    <a @click="clickButoon"
                    href="javascript:void(0)"
                    class="button--grey"
                    >名寄せ実行</a>
                </div>
                <div class="col-3">
                    <h3>名寄せ後</h3>
                    <div class="row folder-list">
                        <div v-for="file in files3" :key="file.id" class="col-12">{{file.Name}}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style>

h3 {
  padding: 0.4em 0.5em;
  color: #494949;
  background: #f4f4f4;
  border-left: solid 5px #7db4e6;
  border-bottom: solid 3px #d7d7d7;
}

.folder-list {
    margin: 0.1rem;
}

.folder-list > div {
    min-height: 50px;
    border-top: 1px solid;
    border-right: 1px solid;
    border-left: 1px solid;
}

.folder-list > div:last-child {
    border-bottom: 1px solid;
}

.folder-list > div:nth-child(even) {
    background: #fff;
}

.folder-list > div:nth-child(odd) {
    background: #ddd;
}

</style>
<script>

export default {

    data: function() {
        return {
            files1: [
                {
                    id: 1,
                    Name: "2018年度会員名簿",
                },
                {
                    id: 1,
                    Name: "2019年度会員名簿",
                },
                {
                    id: 1,
                    Name: "2020年度会員名簿",
                }
            ],
            files2: [
                {
                    id: 1,
                    Name: "雪まつり参加者名簿",
                },
                {
                    id: 1,
                    Name: "センキャン参加者名簿",
                }
            ],
            files3: []
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            console.log('load')
        })
    },
    methods: {
        async clickButoon(){
            this.files3=[
                {
                    id: 1,
                    Name: "名寄せ済み_20200327.xlsx",
                }
            ]
        },
        async clickFile(){
            alert('clickFile')

        }
    }
}

</script>