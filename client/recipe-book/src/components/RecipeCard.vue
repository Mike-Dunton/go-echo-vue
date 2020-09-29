<template>
  <div class="card">
    <div class="card-image">
      <figure class="image is-4by3">
        <img :src="recipe.featuredImage" alt="Placeholder image">
      </figure>
    </div>
    <div class="card-content">
      <div class="media">
        <div class="media-left">
          <figure class="image is-48x48">
            <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
          </figure>
        </div>
        <div class="media-content">
          <p class="title is-4">{{recipe.name}}</p>
          <p class="subtitle is-6">{{recipe.category}}</p>
        </div>
      </div>
      <div class="content">
        {{recipe.description}}
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: [
    'recipe'
  ]
}
</script>