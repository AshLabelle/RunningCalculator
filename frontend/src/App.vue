<template>
  <div class="card">
    <span class="card-title">Calculate your estimated finish time</span>

    <div class="card-content">
      <form v-on:submit.prevent="calcuateTime">
        <label>Distance (km)</label>
        <input
          v-model="distance"
          min="0"
          max="999.99"
          type="number"
          step=".01"
          placeholder="E.g. 10"
          required
          class="form-input"
        />
        <div class="error-text">{{ distanceError }}</div>
        <label>Pace (per km)</label>
        <input
          v-model="pace"
          type="text"
          placeholder="E.g. 6m30s"
          required
          pattern="\d{1,2}m\d{1,2}s"
          class="form-input"
        />
        <div class="error-text">{{ paceError }}</div>
        <input type="submit" value="Calculate!" />
      </form>
      <div>{{ result }}</div>
      <div class="error-text">{{ errorMessage }}</div>
    </div>
  </div>
</template>

<script src="./app.js"></script>
