<template>
  <div class="center">
    <p>
      </p><div class="input-group">
          <span class="input-group-btn">
              <button type="button" class="btn btn-danger btn-number"  data-type="minus" data-field="quant[2]">
                <Icons :icon="'fa-solid fa-circle-minus'"></Icons>
              </button>
          </span>
          <input type="text" name="quant[2]" class="form-control input-number" value="10" min="1" max="100">
          <span class="input-group-btn">
              <button type="button" class="btn btn-success btn-number" data-type="plus" data-field="quant[2]">
                  <Icons :icon="'fa-solid fa-circle-plus'"></Icons>
              </button>
          </span>
      </div>
	<p></p>
</div>  
</template>

<script>
export default {
  name:"PMButton",
  data(){
    return{

    }
  },

}
</script>

<style>
.center{
  width: 150px;
  margin: 40px auto;
  
}
</style>