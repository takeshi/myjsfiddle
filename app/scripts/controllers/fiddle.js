'use strict';

/**
 * @ngdoc function
 * @name myjsfiddleApp.controller:FiddleCtrl
 * @description
 * # FiddleCtrl
 * Controller of the myjsfiddleApp
 */
angular.module('myjsfiddleApp')
  .controller('FiddleCtrl', function ($scope,$state,$stateParams,$sce,ContentsService) {

    var themeId = $stateParams.themeId;
    var contentsId = $stateParams.contentsId;

    var initHtml = '' +
'<div ng-app="sample">\n' +
'  <div ng-repeat="a in [1,2,3] ">\n' +
'    <p>{{a}}</p>\n' +
'  </div>\n' +
'</div>\n'; 
    var initCss = '' +
'p {\n' +
' color:red\n' +
'}\n' +
'';
    var initJs = '' +
"angular\n" +
"  .module('sample', [\n" +
"    'ngAnimate',\n" +
"    'ngCookies',\n" +
"    'ngResource',\n" +
"    'ui.router',\n" +
"    'ngSanitize',\n" +
"    'ngTouch'\n" +
"  ]);"
'';

    if(themeId == 'new'){
      $scope.theme = {
        Contents:{
          Html : initHtml,
          Css: initCss,
          Js : initJs,
        }
      };
    }else{
      ContentsService.find(themeId,contentsId)
      .success(function(theme){
        $scope.theme = theme;
        console.log(theme)
      });
    }
    $scope.currentProjectUrl = $sce.trustAsResourceUrl('/app/run/'+ themeId +'/' + contentsId);

    $scope.run = function(){
      ContentsService.createTheme($scope.theme)
      .success(function(theme){
        console.log(theme);

        $state.go('main.fiddle',{
          themeId:theme.Id,
          contentsId:theme.Contents.Id
        });
      });
    };

  });
