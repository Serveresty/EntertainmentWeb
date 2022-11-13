$(function () {
  
    // Restrict input to decimal points with text masking
    $('input.number').on('blur focus change init-number', function() {
      var $this = $(this),
        data  = $this.data(),
        max   = $this.attr('max'),
        min   = $this.attr('min'),
        val   = this.value.replace(/[^0-9.]/,'') != '' ? (+this.value.replace(/[^0-9.]/g,'')) : '',
        
        numeric = function (n) {
          return !isNaN(n) && isFinite(n);
        };
  
      if (val !== '') {
        numeric(max) && Math.min(val, max);
        numeric(min) && Math.max(val, min);
  
        if (numeric(data.decimal)) { 
          val = val.toFixed(data.decimal);
  
          if(data.trailingZeroes === undefined) {
            val = +val;
          }
        }
      }
  
      this.value = val;
    }).on('input', function() {
      var val   = this.value.replace(/[^0-9.]/,''),
        arr   = val.split('.'),
        first = arr.shift(),
        join  = arr.join('');
  
      this.value = arr.length > 1 ? first+'.'+join : val;
    }).trigger('init-number');
  
    // Mask input as phone number
    // TODO: Retain caret position after input
    $('input.phone').on('blur focus change input init-phone', function(e) {
      var y = this.value.replace(/\W/g, ''),
        x = y.match(/(\w{0,3})(\w{0,3})(\w{0,4})/);
      this.value = (y.length <= 10) ? (!x[2] ? x[1] : '(' + x[1] + ') ' + x[2] + (x[3] ? '-' + x[3] : '')) : ('+' + y);
    }).trigger('init-phone');
  
    // Mask input as date
    // TODO: Retain caret position after input
    // TODO: keyup/keydown increment/decrements year/month/day depending on caret position
    var oldval = false;
    $('input.date').on('blur focus change input init-date', function(e) {
      
      // store current positions in variables
      var start = this.selectionStart,
        value = this.value.replace(/[\D]/g, '').slice(0, 8),
        format = value;
      
      if (value.length > 6) {
        format = value.replace(/(\d{4})(\d{2})(\d*)/, '$1-$2-$3');
      } else if (value.length > 4) {
        format = value.replace(/(\d{4})(\d*)/, '$1-$2');
      }
  
      offset = format.length - Math.min(10, this.value.length);
      start = start+offset;
      
      this.value = format;
      this.selectionStart = start;
      this.selectionEnd = start;
    }).trigger('init-date');  
  });