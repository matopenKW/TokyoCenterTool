<template>
    <div id="top">
        <div class="row">
            <div class="col-3">
                <Menu />
            </div>
            <div class="col-9">
                <div class="row">
                    <div id="carfare" class="col-4 text-center" @click="clickCarfare">
                        <div class="title">交通費管理</div>
                        <div class="icon">
                            <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="subway" class="svg-inline--fa fa-subway fa-w-14" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M448 96v256c0 51.815-61.624 96-130.022 96l62.98 49.721C386.905 502.417 383.562 512 376 512H72c-7.578 0-10.892-9.594-4.957-14.279L130.022 448C61.82 448 0 403.954 0 352V96C0 42.981 64 0 128 0h192c65 0 128 42.981 128 96zM200 232V120c0-13.255-10.745-24-24-24H72c-13.255 0-24 10.745-24 24v112c0 13.255 10.745 24 24 24h104c13.255 0 24-10.745 24-24zm200 0V120c0-13.255-10.745-24-24-24H272c-13.255 0-24 10.745-24 24v112c0 13.255 10.745 24 24 24h104c13.255 0 24-10.745 24-24zm-48 56c-26.51 0-48 21.49-48 48s21.49 48 48 48 48-21.49 48-48-21.49-48-48-48zm-256 0c-26.51 0-48 21.49-48 48s21.49 48 48 48 48-21.49 48-48-21.49-48-48-48z"></path></svg>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    components: {
        Menu: () => import ('~/components/Menu.vue'),
    },
    methods: {
        async clickCarfare(){
            window.location.href = './carfare'
        }
    }
}
</script>

<style>
.row {
    margin: 0 3vh;
}
#carfare {
    margin: 0 1vw;
    margin-bottom: 5.5vh;
    box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
    height: 450px;
    cursor: pointer;
    
}

#carfare .title {
    font-size: 2.5vh;
    margin: 2vh 0;
}

#carfare .icon {
    height: 80%;
    background: #e9ecef;
    vertical-align: middle;
    display: flex;
    align-items: center;
	justify-content: center;

}

#carfare .icon svg {
    height: 150px;
}


</style>