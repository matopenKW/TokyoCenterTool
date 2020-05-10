<template>
    <div id="menu">
        <div class="box">
            <button class="btn" @click="clickBtn">Home<i class="fas fa-home"></i></button>
            <div class="panel">
                <h2>Home</h2>
            </div>
            <button class="btn">About<i class="fas fa-users"></i></button>
            <div class="panel">
                <h2>About</h2>
            </div>
            <button class="btn" >Menu<i class="fas fa-chevron-circle-down"></i></button>
            <div class="panel">
                <h2>Menu</h2>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    mounted: function () {
        },
        methods: {
            async clickBtn(obj){
                console.log(obj)
            }
        }

}
</script>

<style>
#menu .panel {
  border: 1px solid #acac;
  border-left: none;
  border-right:none;
  background-color: white;
  padding: 20px 20px;
  display: none;
}
#menu .box {
  background-color: #eee;
  display: block;
  display: inline-block;
  width: 100%;
  border: 1px solid #acac;
}
#menu .btn{
  display: block;
  flex-grow: 1;
  width: 100%;
  padding: 20px;
  background-color: inherit;
  outline: none;
  border: none;
  cursor: pointer;
  text-align: left;
}
#menu .btn:hover {
  background-color: #ddd;
}
#menu .fas {
  float: right;
}
</style>