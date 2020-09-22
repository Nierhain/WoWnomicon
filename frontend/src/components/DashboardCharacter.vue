<template>
  
</template>

<script>
export default {
    name: "DashboardCharacter",
    props: {
        character: Object
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
        this.character.thumbnail
      );
    },
    itemLevel: function() {
      return this.character.items.averageItemLevelEquipped;
    },
    guildName: function() {
      return this.character.guild.name;
    },
    charClass: function() {
      return classToString(this.character.class);
    }
    }
}
</script>

<style>

</style>