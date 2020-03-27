<template>
    <div>
        test2
    </div>
</template>

<script>
import RUOTES from '~/routes/api'

export default {
    async fetch({store}) {
        const payload = {
            uri: RUOTES.GET.PROFILE
        }

        await store.dispatch("test",payload)
    }
}

</script>