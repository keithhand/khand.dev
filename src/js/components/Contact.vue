<template>
  <section id="contact">
    <h4><strong>{{ name }}</strong></h4>
    <div>
      <table>
        <tr>
          <font-awesome-icon :icon="['fa','phone']" />
          {{ phone }}
        </tr>
        <tr>
          <font-awesome-icon :icon="['fas','envelope']" />
          <a>
            {{ email }}
          </a>
        </tr>
        <tr>
          <font-awesome-layers
            class="fa-3x"
            style="padding-top: 50px"
          >
            <font-awesome-icon :icon="['fab', 'github']" />
          </font-awesome-layers>
        </tr>
      </table>
    </div>
  </section>
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      required: true,
    },
    phone: {
      type: Number,
      required: true,
    },
    email: {
      type: String,
      required: true,
    },
  },

  data() {
    return {
      hostname: window.location.hostname,
    };
  },
};
</script>

<style lang="scss" scoped>
  @import "../../assets/css/variables";

  section {
    background: $main-color;
    color: white;
    font-size: 14px;
    padding: 50px 0 10px 0;

    div {
      display: flex;
      align-items: center;
      justify-content: center;
      margin: auto;
    }

    a {
      color: white;
    }
  }
</style>
