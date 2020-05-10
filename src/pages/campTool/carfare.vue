<template>
    <div id="carfare">
        <div class="row">
            <div class="menu col-3">
                <Menu />
            </div>
            <div id="carfare-area" class="col-sm-12 col-md-9 col-lg-9">
                <h2>交通費管理</h2>
                <div id="carfare-serch-area" class="row">
                    <div class="col-12">
                        検索機能 Comming Soon...
                    </div>
                </div>
                <div id="carfare-list-area" class="row">
                    <table>
                        <thead>
                            <tr>
                                <td class="date-head">日付</td>
                                <th scope="col">内容</th>
                                <th class="price-head" scope="col">価格(月額)</th>
                                <th class="btn-area-head" scope="col"></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="carfare in carfareList" :key="carfare.SeqNo">
                                <th class="date">{{ carfare.Date }}</th>
                                <td data-label="内容" class="route">
                                    <div class="row">
                                        <div class="col-md-12 col-lg-5">
                                            <input type="text" v-model="carfare.Start">
                                        </div>
                                        <div class="arrow text-center col-md-12 col-lg-1"></div>
                                        <div class="col-md-12 col-lg-5">
                                            <input type="text" v-model="carfare.End">
                                        </div>
                                    </div>
                                </td>
                                <td data-label="価格" class="price">{{ carfare.Price }}</td>
                                <td data-label="" class="btn-area">
                                    <input type="button" class="btn btn-primary" value="登録" click="clickRegistBtn">
                                    <input type="button" class="btn btn-danger" value="削除">
                                </td>
                            </tr>
                        </tbody>
                    </table>   
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    components: {
        Menu: () => import ('~/components/Menu.vue'),
    },
    data: function(){
        return {
            carfareList: []
        }
    },
    mounted: function(){
        this.$nextTick(function () {
            this.$axios.$get('/getCarfare',{
                params:{}
            })
            .then(response => {
                this.carfareList = response
            })
            .catch(err => {
                console.log(err)
            })
            .finally(() => console.log('finally'))
        })
    },
    methods: {
        async clickRegistBtn(){
            console.log(this.carfareList)
        }
    }
}
</script>

<style>

#carfare-area h2 {
    margin-bottom: 2vh;
}

#carfare-area #carfare-serch-area {
    margin-bottom: 2vh;
    height: 50px;
    background-color: #eee;
    border: 1px solid #bbb;
}

#carfare-area #carfare-serch-area > div {
    margin: auto;
}

#carfare-area #carfare-list-area {
    justify-content: center;
}

#carfare-area #carfare-list-area .btn {
    width: 76px;
}

@media screen and (max-width: 768px) {
    .menu {
        display: none;
    }
}

/*テーブルレイアウト */
table .date-head {
    width: 110px;
}

table .price-head {
    width: 110px;
}

table .btn-area-head {
    width: 200px;
}

table .route input{
    width: 100%;
}

table .route .arrow:after {
    content: '→';
}

@media screen and (max-width: 992px) {
    table .route .arrow:after {
        content: '↓';
    }
}

/*テーブルレイアウト defalt*/
table {
  border-collapse: collapse;
  margin: 0 auto;
  padding: 0;
  width: 100%;
  table-layout: fixed;
}

table tr {
  background-color: #fff;
  border: 1px solid #bbb;
  padding: .35em;
}
table th,
table td {
  padding: 1em 10px 1em 1em;
  border-right: 1px solid #bbb;
}
table th {
  font-size: .85em;
}
table thead tr{
  background-color: #eee;
}
table .route{
   text-align: left;
   font-size: .85em;
}
table .price{
   text-align: right;
}
@media screen and (max-width: 768px) {

  table {
    border: 0;
    width:100%
  }
  table th{
    background-color: #eee;
    display: block;
    border-right: none;
  }
  table thead {
    border: none;
    clip: rect(0 0 0 0);
    height: 1px;
    margin: -1px;
    overflow: hidden;
    padding: 0;
    position: absolute;
    width: 1px;
  }
  
  table tr {
    display: block;
    margin-bottom: .625em;
  }
  
  table td {
    border-bottom: 1px solid #bbb;
    display: block;
    font-size: .8em;
    text-align: right;
    position: relative;
    padding: .625em .625em .625em 4em;
    border-right: none;
  }
  
  table td::before {
    content: attr(data-label);
    font-weight: bold;
    position: absolute;
    left: 10px;
  }
  
  table td:last-child {
    border-bottom: 0;
  }
}
</style>