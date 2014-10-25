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
    'ngTouch'
  ])
  .config(function ($stateProvider,$urlRouterProvider) {
    $urlRouterProvider.otherwise('/fiddle/new/0');

    $stateProvider
      .state('main', {
        url: '',
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      
      .state('main.fiddle',{
        url: '/fiddle/:themeId/:contentsId',
        templateUrl: 'views/fiddle.html',
        controller: 'FiddleCtrl'        
      })
      ;

  });
