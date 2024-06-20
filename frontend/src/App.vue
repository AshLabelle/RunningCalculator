<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Calculate your estimated finish time</h1>

        <form v-on:submit.prevent="calcuateTime">
          <div class="form-group">
            <input
              v-model="distance"
              type="text"
              id="website-input"
              placeholder="Enter a distance in km"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <input
              v-model="pace"
              type="text"
              id="website-input"
              placeholder="Enter a pace (mm:ss) per km"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Calculate!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",

  data() {
    return {
      distance: 10,
      pace: "6:00",
    };
  },

  methods: {
    calcuateTime() {
      console.log(
        `Should calculate the finish time of ${this.distance}km at ${this.pace}/km avg pace`
      );
      var postData = {
        distance: this.distance,
        pace: this.pace,
      };

      let axiosConfig = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      // Call the Go API, in this case we only need the URL parameter.
      axios
        .post("http://localhost:3000/api/calculateTime", postData, axiosConfig)
        .then((response) => {
          console.log("response:");
          console.log(response.data);
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
  },
};
</script>
