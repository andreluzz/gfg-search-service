<template>
  <div id="app">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Products
          </h1>
          <h2 class="subtitle">
            Search products filtering, sorting and paginating
          </h2>
        </div>
      </div>
    </section>
    <div class="container">
      <Products/>
    </div>
  </div>
  
</template>

<script>
import Products from './components/Products.vue'

export default {
  name: 'app',
  components: {
    Products
  }
}
</script>