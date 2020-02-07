<template>
  <section id="languages">
    <h2>My Knowledge Base</h2>
    <hr>
    <ul>
      <div id="language-list">
        <li
          v-for="l in languages"
          :key="l.name"
          :style="{'width': numberOfRows}"
        >
          <div class="fontawesome-icon">
            <font-awesome-layers class="fa-4x fa-stack">
              <font-awesome-icon
                :icon="['fa', 'circle']"
                class="fa-stack-2x"
              />
              <font-awesome-icon
                :icon="l.icon"
                class="fa-stack-1x"
                style="color: #333"
              />
            </font-awesome-layers>
          </div>
          {{ l.title }}
        </li>
      </div>
    </ul>
  </section>
</template>

<script>
export default {
  props: {
    languages: {
      type: Array,
      required: true,
    },
  },

  computed: {
    numberOfRows() {
      const { length } = this.languages;
      if (length % 2 === 0 && length % 3 !== 0) {
        return '20%';
      }
      return '30%';
    },
  },
};
</script>

<style scoped lang="scss">
  @import "../../assets/css/variables";

  section {
    color: white;
    background: $main-color;
  }

  .fontawesome-icon {
    transform: scale(.8);
    margin-top: 10px;
  }

  #language-list {
    margin: 0 100px;
  }
</style>
