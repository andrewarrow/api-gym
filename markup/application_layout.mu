html attr-1
  head
    {{ $build := index . "build" }}
    {{ $og := index . "og" }}
    meta property=og:image content={{$og}}
    link rel=apple-touch-icon href=/assets/images/ls.png
    link rel=apple-touch-startup-image href=/assets/images/ls.png
    link rel=icon href=/assets/images/ls.png
    link rel=stylesheet type=text/css href=/assets/css/tail.min.css?id!{{$build}}
    link rel=stylesheet type=text/css href=https://cdnjs.cloudflare.com/ajax/libs/noUiSlider/15.7.1/nouislider.css
    link rel=stylesheet type=text/css href=/assets/css/main.css?id!{{$build}}
    {{ if index . "USE_LIVE_TEMPLATES" }}
      script src=https://cdn.tailwindcss.com
    {{ end }}
    script src=/assets/javascript/wasm_exec.js?id!{{$build}}
    script
      function $(id) { return document.getElementById(id); }
    script src=https://cdnjs.cloudflare.com/ajax/libs/noUiSlider/15.7.1/nouislider.min.js
    title
      {{ index . "title" }}
    {{ index . "viewport" }}
  body
    div id=flash bg-red-600 text-white text-center fixed top-0 left-0 w-full
      {{ index . "flash" }}
    div bg-black text-gray-300 overflow-x-auto pl-3 pr-3 min-h-screen font-poppins text-base
      {{ index . "content" }}
    {{ index . "wasm" }}
