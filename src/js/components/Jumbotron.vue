<template>
  <section
    id="jumbotron"
    class="parallax-container"
  >
    <parallax
      direction="down"
      :speed-factor=".25"
    >
      <img
        src="../../assets/img/bg.jpg"
        alt="homepage background"
        class="image"
      >

      <h1 class="parallax-title">
        {{ name }}
      </h1>
      <h3 class="parallax-subtitle">
        {{ title }}
      </h3>
    </parallax>
    <!--    <a>Who am I?</a>-->
  </section>
</template>

<script>
import Parallax from 'vue-parallaxy';

export default {
  components: {
    Parallax,
  },

  props: {
    name: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
  },
};
</script>

<style lang="scss" scoped>
  .parallax-container {
    position: relative;
    height: 1000px;
  }

  .parallax-title {
    position: absolute;
    top: 30%;
    left: 0;
    right: 0;
    padding: 20px;
    font-size: 4.5em;
    font-weight: 700;
    color: white;
    text-align: center;
  }

  .parallax-subtitle {
    position: absolute;
    top: 40%;
    left: 0;
    right: 0;
    padding: 20px;
    font-size: 2.25em;
    font-weight: 600;
    color: white;
    text-align: center;
  }

  .image {
    height: 700px;
  }

  section {
    padding: 0;
  }
</style>
