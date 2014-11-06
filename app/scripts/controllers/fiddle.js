'use strict';

/**
 * @ngdoc function
 * @name myjsfiddleApp.controller:FiddleCtrl
 * @description
 * # FiddleCtrl
 * Controller of the myjsfiddleApp
 */
angular.module('myjsfiddleApp')
  .controller('FiddleCtrl', function ($scope,$state,$stateParams,$sce,ContentsService,$window) {

    var themeId = $stateParams.themeId;
    var contentsId = $stateParams.contentsId;
    var mode = $stateParams.mode;

    $scope.Mode.perspective = mode;

    var initHtml = '' +
'<div ng-controller="sampleCtrl" class="container">\n' +
'  <div ng-repeat="item in items ">\n' +
'    <p>{{item}}</p>\n' +
'  </div>\n' +
'  <div sample-directive="" ></div>' +
'</div>\n'; 
    var initCss = '' +
'div{\n'+
'  p {\n' +
'   color:red\n' +
'  }\n'+
'}\n' +
'';
    var initJs = '' +
"angular\n" +
"  .module('sample', [\n" +
"  ]);\n" +
" \n" +
"angular\n" +
" .module('sample')\n" +
" .controller('sampleCtrl',function($scope){\n" +
"    $scope.items = [1,2,3];\n" +
" })\n" +
" \n" +
"angular\n" +
" .module('sample')\n" +
" .directive('sampleDirective',function(){\n" +
"    return {\n" +
"      templateUrl:'/directive.html' \n" +
"    }" +
" })\n" +
';';

    var initDirective = ''+
"<div> Hello Directive </div>"
'';

    if(themeId == null || themeId == 'new'){
      $scope.theme = {
        Contents:{
          Html : initHtml,
          Css: initCss,
          Js : initJs,
          Directive : initDirective
        }
      };
    }else{
      ContentsService.find(themeId,contentsId)
      .success(function(theme){
        $scope.theme = theme;
        console.log(theme)
      });
    }
    $scope.currentProjectUrl = $sce.trustAsResourceUrl('/fiddle/app/run/'+ themeId +'/' + contentsId);

    $scope.run = function(){
      ContentsService.createTheme($scope.theme)
      .success(function(theme){
        console.log(theme);

        $state.go('main.fiddle',{
          themeId:theme.Id,
          contentsId:theme.Contents.Id,
          mode:$scope.Mode.perspective
        });
      });
    };

    $scope.open = function(){
      $window.location.href = '/fiddle/app/run/'+ themeId +'/' + contentsId;
    }
  });
