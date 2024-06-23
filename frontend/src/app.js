import "./app.css";
import axios from "axios";

export default {
  name: "App",

  data() {
    return {
      distance: "",
      pace: "",
      result: "",
      distanceError: "",
      paceError: "",
      errorMessage: "",
    };
  },

  methods: {
    calcuateTime() {
      let distanceErrorString = "";
      let paceErrorString = "";

      if (distanceErrorString === "" && paceErrorString === "") {
        let postData = {
          distance: this.distance,
          pace: this.pace,
        };

        let axiosConfig = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        console.log("Request:");
        console.log(postData);

        // Call the Go API, in this case we only need the URL parameter.
        axios
          .post(
            "http://localhost:3000/api/calculateTime",
            postData,
            axiosConfig
          )
          .then((response) => {
            this.result = `Your estimated finish time is ${response.data.time}`;
            this.errorMessage = "";
            console.log("Response:");
            console.log(response.data);
          })
          .catch((error) => {
            this.result = "";
            this.errorMessage = error.response.data;
            // window.alert(`The API returned an error: ${error}`);
          });
      }
      this.distanceError = distanceErrorString;
      this.paceError = paceErrorString;
    },
  },
};
