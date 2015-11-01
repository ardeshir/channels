javascript:void( function(){

	    function configRequireJS(){

	   	requirejs.config({  

	   		        paths: {

	   		   	         parser:'https://s3.amazonaws.com/qwic/app/parser',
	   		   	         jquery:'https://s3.amazonaws.com/qwic/vendor/jquery/jquery-2.1.3.min'
	   		   	     },
	   		   	         deps:['parser']
	   		   	    });

	    }

	   function getRequireJs(){

	   	   var s = document.createElement("script");

	   	   s.src="https://s3.amazonaws.com/qwic/vendor/require/require.js";

	   	   if(s.addEventListener){ s.addEventListener("load",configRequireJS,false)

	   	    } else if(s.readyState){
	   	    	s.onreadystatechange=configRequireJS
	   	    }

	   	  document.body.appendChild(s);}

	   getRequireJs();

})();
/*
javascript:void(function (){

       function configRequireJS(){

            requirejs.config({

                   	  paths:{
                                 parser:'http://api.metalearn.org/thanks',
                                 jquery:'https://s3.amazonaws.com/qwic/vendor/jquery/jquery-2.1.3.min'
                            },
                       deps:['thanks']});

        }

        function getRequireJs(){

               var s=document.createElement("script");

               s.src="http://api.metalearn.org/require.js";

             if(s.addEventListener){
             	s.addEventListener("load",configRequireJS,false) 
             } else if(s.readyState) {
             	s.onreadystatechange=configRequireJS
             } 
             document.body.appendChild(s);
        }
     getRequireJs();

 })(); */