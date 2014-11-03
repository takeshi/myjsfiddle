'use strict';

/**
 * @ngdoc overview
 * @name myjsfiddleApp
 * @description
 * # myjsfiddleApp
 *
 * Main module of the application.
 */
angular
  .module('myjsfiddleApp', [
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ui.router',
    'ngSanitize',
    'ngTouch',
    'ui.codemirror'
  ])
  .config(function ($stateProvider,$urlRouterProvider) {
    $urlRouterProvider.otherwise('/fiddle/new/0/css');

    $stateProvider
      .state('main', {
        url: '',
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      
      .state('main.fiddle',{
        url: '/fiddle/:themeId/:contentsId/:mode',
        templateUrl: 'views/fiddle.html',
        controller: 'FiddleCtrl'        
      })
      ;

  });
