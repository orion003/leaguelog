angular.module('rlApp').config(['$stateProvider', '$urlRouterProvider', '$locationProvider',
    function($stateProvider, $urlRouterProvider, $locationProvider) {
        $stateProvider
            .state('home', {
                url: '/',
                views: {
                    '': {
                        templateUrl: '/app/components/home/home.html'
                    },
                    'search@home': {
                        templateUrl: '/app/components/home/home.search.html',
                        controller: 'HomeSearchController',
                        controllerAs: 'home'
                    },
                    'register@home': {
                        templateUrl: '/app/components/home/home.register.html',
                        controller: 'HomeRegisterController',
                        controllerAs: 'home'
                    }
                }
            })
            .state('league', {
                'url': '/l/{leagueId:int}',
                views: {
                    '': {
                        templateUrl: '/app/components/league/league.html'
                    },
                    'header@league': {
                        templateUrl: '/app/components/league/league.header.html',
                        controller: 'LeagueHeaderController',
                        controllerAs: 'league'
                    },
                    'standings@league': {
                        templateUrl: '/app/components/league/league.standings.html',
                        controller: 'LeagueStandingsController',
                        controllerAs: 'league'
                    },
                    'schedule@league': {
                        templateUrl: '/app/components/league/league.schedule.html',
                        controller: 'LeagueScheduleController',
                        controllerAs: 'league'
                    }
                }
            })
            .state('screen', {
                url: '/screen',
                templateUrl: '/app/components/screen/screen.html'
            });
            
        $urlRouterProvider.otherwise('/');
        
        $locationProvider.html5Mode(true);
    }]);