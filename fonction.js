
monStockage = localStorage;



function resetParam() {
    localStorage.setItem('monID', undefined);
    localStorage.setItem('monNom', undefined);
    localStorage.setItem('bat_1', undefined);
  }

$(document).ready(function() {

    if (monStockage.monNom != undefined) {
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/joueur/liste/1",
            success: function (response) {
                if (response.Data == monStockage.monNom){
                    monID = 1;
                }else {
                    monID = 2;
                }
                localStorage.setItem('monID', monID);
                console.log(monID);
            },
            error: function (params) {
            //   console.log(params)  
            },
            done: function () {
                // console.log("test3");
            }
            
        });
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/jeux/carte/"+monStockage.monID,
            success: function (response) {
                refreshCarte(response.Data)
            },
            error: function (params) {
            //   console.log(params)  
            },
            done: function () {
                // console.log("test3");
            }
            
        });



        

        setInterval(function()
        { 
            $.ajax({
            type:"post",
            url:"http://localhost:56700/jeux/tour",
            datatype:"html",
            success:function(data)
            {
                if (data.Data == monStockage.monID && monStockage.bat_1 == "O") {
                    $("#zone_tir").show();
                    $.ajax({
                        type: "GET",
                        url: "http://localhost:56700/jeux/carte/"+monStockage.monID,
                        success: function (response) {
                            refreshCarte(response.Data)
                        },
                        error: function (params) {
                        //   console.log(params)  
                        },
                        done: function () {
                            // console.log("test3");
                        }
                        
                    });
                }else {
                    $("#zone_tir").hide();
                    $.ajax({
                        type: "GET",
                        url: "http://localhost:56700/jeux/carte/"+monStockage.monID,
                        success: function (response) {
                            refreshCarte(response.Data)
                        },
                        error: function (params) {
                        //   console.log(params)  
                        },
                        done: function () {
                            // console.log("test3");
                        }
                        
                    });
                }
            }
            });
        }, 1000);//time in milliseconds 

        if (monStockage.bat_1 == "O") {
            $("#zone_bat_1").hide();
            
        }else{
            $("#zone_tir").hide();

        }

    }


    $("#liste").click(function (e) {
        console.log("click liste");
        e.preventDefault();
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/joueur/liste",
            success: function (response) {
                // console.log("test2");
                // console.log(response.Msg);
                $("#Msg").text(response.Msg);
                $("#data").text(JSON.stringify(response.Data));
            },
            error: function (params) {
            //   console.log(params)  
            },
            done: function () {
                // console.log("test3");
            }
            
        });    
    });
    $("#fifi").click(function (e) {
        e.preventDefault();
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/joueur/name/fifi",
            success: function (response) {
                $("#Msg").text(response.Msg);
                $("#data").text(JSON.stringify(response.Data));
            },
            error: function (params) {
            //   console.log(params)  
            },
            done: function () {
                // console.log("test3");
            }
            
        });    
    });
    $("#riri").click(function (e) {
        e.preventDefault();
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/joueur/name/riri",
            success: function (response) {
                $("#Msg").text(response.Msg);
                $("#data").text(JSON.stringify(response.Data));
                
            },
            error: function (params) {
            //   console.log(params)  
            },
            done: function () {
                // console.log("test3");
            }
            
        });    
        console.log(monNom)
    });
    $("#ValideNom").click(function (e) {
        e.preventDefault();
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/joueur/name/"+$("#inp_nom").val(),
            success: function (response) {
                if (response.Ok == true) {
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                    localStorage.setItem('monNom', response.Data);
                    $("#ValideNom").hide();
                    $("#carte").show();
                }else{
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                }
            },
            error: function (params) {
              console.log(params)  
            },
            done: function () {
                console.log("test3");
            }
        });    
        
    });
    $("#valid_bat_1").click(function (e) {
        e.preventDefault();

        console.log("http://localhost:56700/jeux/placer/"+monStockage.monID+"/"+$("#val_empl_bat_1").val());

        $.ajax({
            type: "GET",
            url: "http://localhost:56700/jeux/placer/"+monStockage.monID+"/"+$("#val_empl_bat_1").val(),
            success: function (response) {
                if (response.Ok == true) {
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                    localStorage.setItem("bat_1","O");
                    $("#zone_bat_1").hide();
                    refreshCarte(response.Data)
                }else{
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                }
            },
            error: function (params) {
              console.log(params)  
            },
            done: function () {
                console.log("test3");
            }
        });    
        
    });
    $("#valid_tir").click(function (e) {
        e.preventDefault();
        console.log("http://localhost:56700/jeux/tir/"+$("#val_tir").val());
        $.ajax({
            type: "GET",
            url: "http://localhost:56700/jeux/tir/"+$("#val_tir").val(),
            success: function (response) {
                if (response.Ok == true) {
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                    refreshCarte(response.Data)
                }else{
                    $("#Msg").text(response.Msg);
                    $("#data").text(response.Data)
                }
            },
            error: function (params) {
              console.log(params)  
            },
            done: function () {
                console.log("test3");
            }
        });    
        
    });
    $("td").click(function (e) {
        idParent = $(this).parent().attr('id');
        idThis = $(this).attr('id');

        if (monStockage.bat_1 == "O") {
            $("td").css("background-color", "white");
            $("tr[id="+idParent+"] > td[id="+idThis+"] ").css("background-color", "yellow");
            $("#val_tir").val(idParent+idThis)
        }else{

            console.log($(this).parent().attr('id')+$(this).attr('id'));
            $("td").css("background-color", "white");
            if ($( "#ori_1 option:selected" ).val() == "H") {
                // $(this).css("background-color", "red");
                $("tr[id="+idParent+"] > td[id="+idThis+"] ").css("background-color", "red");
                $("tr[id="+idParent+"] > td[id="+(Number(idThis)+1)+"] ").css("background-color", "red");
                $("tr[id="+idParent+"] > td[id="+(Number(idThis)+2)+"] ").css("background-color", "red");

            }else{
                $("tr[id="+idParent+"] > td[id="+idThis+"] ").css("background-color", "red");
                $("tr[id="+(Number(idParent)+1)+"] > td[id="+idThis+"] ").css("background-color", "red");
                $("tr[id="+(Number(idParent)+2)+"] > td[id="+idThis+"] ").css("background-color", "red");
            }
            $("#val_empl_bat_1").val($( "#ori_1 option:selected" ).val()+"/"+$(this).parent().attr('id')+$(this).attr('id'))
        }
    });
});

function refreshCarte(param) { 
    for (x in param) {
        for (y in param[x]) {
            // console.log(x+y)
            if (param[x][y] == monStockage.monID) {
                $("tr[id="+x+"] > td[id="+y+"] ").css("background-color", "blue");
            }else if (param[x][y] == 3) {
                $("tr[id="+x+"] > td[id="+y+"] ").css("background-color", "black");
            }
        }
    }
 }


// console.log("test4");