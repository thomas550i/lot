import ClientGLobalComponents from "./ClientGlobalComponets";
import "../assets/css/bootstrap.min.css";
import "../css/theme.min.css";
import "../assets/fonts/icofont/icofont.min.css";
import "../assets/fonts/icofont/icofont.min.css"
import "../assets/js/jquery.min.js"
// import "../assets/js/bootstrap.min.js"
import "../assets/js/jquery-ui.min.js"
// import "../assets/js/owl.carousel.min.js"
// import "../assets/js/lv-ripple.jquery.min.js"
import "../assets/js/SmoothScroll.min.js"
// import "../assets/js/jquery.TDPageEvents.min.js"
// import "../assets/js/jquery.TDParallax.min.js"
// import "../assets/js/jquery.TDTimer.min.js"
import "../assets/js/selectize.min.js"
// import "../js/main.min.js"


const RequiredComponets={
  install(app){
    app.component("ClientGLobalComponents", ClientGLobalComponents);
  }
}
export default RequiredComponets;
