<template>
    <div>
        <section id="profile">
            <profile />
        </section>
        <section id="skill">
            <skill />
        </section>
        <section id="product">
            <product />
        </section>
    </div>
</template>

<script>

export default {
    components: {
        profile: () => import ('~/components/Profile.vue'),
        skill: () => import ('~/components/Skill.vue'),
        product: () => import ('~/components/Product.vue')
    }
}

</script>