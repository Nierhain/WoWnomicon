<template>
  <v-card>
    <v-card-title>{{ characterName }}</v-card-title>
  </v-card>
</template>

<script>
export default {
    name: "DashboardCharacter",
    props: {
        character: Array
    },
    computed: {
    characterName: function () {
      return this.character.name;
    },
    level: function() {
      return this.character.level;
    },
    avatarUrl: function() {
      return (
        "http://render-eu.worldofwarcraft.com/character/" +
        this.character.avatarURL
      );
    },
    itemLevel: function() {
      return this.character.itemlevel;
    },
    guild: function() {
      return this.character.guild;
    },
    className: function() {
      return this.character.className;
    }
    }
}
</script>

<style>

</style>