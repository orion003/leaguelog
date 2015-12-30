angular.module("home", []);
angular.module("league", []);

angular.module('rlApp', [
    'ui.router',
    'angularytics',
    'home',
    'league']);;angular.module('rlApp')
  .config(function(AngularyticsProvider) {
    AngularyticsProvider.setEventHandlers(['GoogleUniversal']);
  }).run(function(Angularytics) {
    Angularytics.init();
  });;angular.module('rlApp').config(['$stateProvider', '$urlRouterProvider', '$locationProvider',
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
    }]);;angular.module('home')
    .controller('HomeSearchController', ['$state', 'HomeService',
        function($state, HomeService){
            ctrl = this;
            HomeService.getLeagues().then(function(result) {
                ctrl.leagues = result.data;
            });
        }])
    .controller('HomeRegisterController', ['$state', 'HomeService',
        function($state, HomeService) {
            this.saveUser = function() {
                HomeService.saveUser(this.user).then(
                    function(response) {
                        $("#registration-form").fadeOut("fast", function() {
                            $(this).html('<div class="row"><div class="small-12 columns email-success">Email submitted. Thank you!</div></div>');
                            $(this).fadeIn("slow");
                        });
                    }, 
                    function(response) {
                        console.log(response);
                        if(response.data.error == "user_duplicate_email") {
                            $("#registration-form").fadeOut("fast", function() {
                                $(this).html('<div class="row"><div class="small-12 columns email-success">Email previously submitted. Thank you!</div></div>');
                                $(this).fadeIn("slow");
                            });
                        }
                        else {
                            $("#registration-form").fadeOut("fast", function() {
                                $(this).html('<div class="row"><div class="small-12 columns email-fail">There was an error. Please refresh and try again.</div></div>');
                                $(this).fadeIn("slow");
                            });
                        }
                    });
            };
        }]);;angular.module('home')
    .service('HomeService', ['$http', function($http) {
            
            this.getLeagues = function() {
                return $http.get('/api/leagues');
            };
            
            this.saveUser = function(user) {
                return $http.post('/api/users', user);  
            };
            
    }]);;angular.module('league')
    .controller('LeagueHeaderController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService){
            var headerController = this;
            LeagueService.getLeague($stateParams.leagueId).then(function(result) {
                headerController.league = result.data;
            });
        }])
    .controller('LeagueStandingsController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService){
            var standingsController = this;
            LeagueService.getStandings($stateParams.leagueId).then(function(result) {
                standingsController.standings = result.data;
            });
        }])
    .controller('LeagueScheduleController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService) {
            var scheduleController = this;
            LeagueService.getSchedules($stateParams.leagueId).then(function(result) {
                scheduleController.schedules = result.data;
            });
        }]);;angular.module('league')
    .service('LeagueService', ['$http', function($http) {
            
            this.getLeague = function(leagueId) {
                return $http.get('/api/league/' + leagueId);
            };
            
            this.getStandings = function(leagueId) {
                return $http.get('/api/league/' + leagueId + '/standings');
            };
            
            this.getSchedules = function(leagueId) {
                return $http.get('/api/league/' + leagueId + '/schedule');
            };
            
    }]);;angular.module('templates-dist', ['../app/components/home/home.html', '../app/components/home/home.register.html', '../app/components/home/home.search.html', '../app/components/league/league.header.html', '../app/components/league/league.html', '../app/components/league/league.schedule.html', '../app/components/league/league.standings.html', '../app/components/screen/screen.html']);

angular.module("../app/components/home/home.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/home/home.html",
    "<header>\n" +
    "    <div class=\"row\">\n" +
    "        <div class=\"small-12 small-centered text-center columns\">\n" +
    "            <div class=\"home-header\">leaguelog</div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "    <div class=\"row\">\n" +
    "        <div class=\"small-12 small-centered text-center columns\">\n" +
    "            <div class=\"home-tagline\">Be a better commish.</div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</header>\n" +
    "<section>\n" +
    "    <div class=\"row small-collapse\">\n" +
    "        <div class=\"small-12 medium-6 columns\">\n" +
    "            <div class=\"home-search\">\n" +
    "                <div data-ui-view=\"search\"></div>\n" +
    "            </div>\n" +
    "        </div>\n" +
    "        <div class=\"small-12 medium-6 columns\">\n" +
    "            <div class=\"home-register\">\n" +
    "                <div data-ui-view=\"register\"></div>\n" +
    "            </div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</section>");
}]);

angular.module("../app/components/home/home.register.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/home/home.register.html",
    "<div class=\"row\">\n" +
    "    <div class=\"small-12 columns\">\n" +
    "        <div class=\"home-section-header\">\n" +
    "            register\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</div>\n" +
    "<div>\n" +
    "    <div id=\"registration-form\">\n" +
    "        <form name=\"home.form\" novalidate>\n" +
    "            <div class=\"row uncollapse\">\n" +
    "                <div class=\"small-12 small-centered columns\">\n" +
    "                    <div class=\"row collapse\">\n" +
    "                        <div class=\"small-12 large-9 columns\">\n" +
    "                            <input type=\"email\" placeholder=\"enter your email\" data-ng-model=\"home.user.email\" required />\n" +
    "                        </div>\n" +
    "                        <div class=\"small-6 small-centered large-3 large-uncentered text-center columns\">\n" +
    "                            <input type=\"submit\" data-ng-click=\"home.saveUser(home.user)\" class=\"button large-postfix\" value=\"submit\" />\n" +
    "                        </div>\n" +
    "                    </div>\n" +
    "                </div>\n" +
    "            </div>\n" +
    "        </form>\n" +
    "    </div>\n" +
    "</div>");
}]);

angular.module("../app/components/home/home.search.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/home/home.search.html",
    "<div class=\"row\">\n" +
    "    <div class=\"small-12 columns\">\n" +
    "        <div class=\"home-section-header\">\n" +
    "            leagues\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</div>\n" +
    "<div class=\"row\">\n" +
    "    <div class=\"small-12 columns\">\n" +
    "        <div class=\"league-list\">\n" +
    "            <ul class=\"no-bullet\">\n" +
    "                <li data-ng-repeat=\"league in home.leagues\">\n" +
    "                    <a data-ui-sref=\"league({leagueId:league.model.id})\" class=\"{{league.sport}}\">{{league.name}}</a>\n" +
    "                </li>\n" +
    "            </ul>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</div>");
}]);

angular.module("../app/components/league/league.header.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/league/league.header.html",
    "<h3 class=\"league-name\">{{league.league.name}}</h3>");
}]);

angular.module("../app/components/league/league.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/league/league.html",
    "<header class=\"main-header-row\">\n" +
    "    <div class=\"row small-uncollapse medium-collapse\">\n" +
    "        <div class=\"small-12 small-centered small-only-text-center medium-uncentered columns\">\n" +
    "            <a data-ui-sref=\"home\"><h1 class=\"main-header\">leaguelog</h1></a>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "    <div class=\"row small-uncollapse medium-collapse\">\n" +
    "        <div class=\"small-12 small-centered medium-uncentered small-only-text-center columns\">\n" +
    "            <h2 class=\"main-tagline\">Be a better commish.</h2>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</header>\n" +
    "<section>\n" +
    "    <div class=\"row small-uncollapse medium-collapse\">\n" +
    "        <div class=\"small-12 small-centered small-only-text-center medium-uncentered columns\">\n" +
    "            <div data-ui-view=\"header\"></div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</section>\n" +
    "<section>\n" +
    "    <div class=\"row medium-collapse\">\n" +
    "        <div class=\"small-12 small-centered medium-4 medium-uncentered columns\">\n" +
    "            <div class=\"main-standings\">\n" +
    "                <div data-ui-view=\"standings\"></div>\n" +
    "            </div>\n" +
    "        </div>\n" +
    "        <div class=\"small-12 medium-8 columns\">\n" +
    "            <div class=\"main-schedule\">\n" +
    "                <div data-ui-view=\"schedule\"></div>\n" +
    "            </div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</section>");
}]);

angular.module("../app/components/league/league.schedule.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/league/league.schedule.html",
    "<div class=\"schedule-list\">\n" +
    "    <div data-ng-repeat=\"schedule in league.schedules | orderBy:'start_date'\">\n" +
    "        <div class=\"schedule-date-group\">\n" +
    "            <div class=\"row medium-collapse\">\n" +
    "                <div class=\"small-12 columns schedule-date\">\n" +
    "                    {{schedule.start_date | date:'fullDate':'UTC'}}\n" +
    "                </div>\n" +
    "            </div>\n" +
    "            <div data-ng-repeat=\"game in schedule.games | orderBy:'game.start_time'\">\n" +
    "                <div class=\"schedule-game\">\n" +
    "                    <div class=\"row\">\n" +
    "                        <div class=\"small-12 medium-2 columns schedule-game-time\">\n" +
    "                            {{game.start_time | date:'shortTime':'UTC'}}\n" +
    "                        </div>\n" +
    "                        <div class=\"small-5 medium-4 columns text-right schedule-game-team\">\n" +
    "                            {{game.away_team.name}}\n" +
    "                        </div>\n" +
    "                        <div class=\"small-2 text-center columns schedule-game-vs\">\n" +
    "                            -vs-\n" +
    "                        </div>\n" +
    "                        <div class=\"small-5 medium-4 columns schedule-game-team\">\n" +
    "                            {{game.home_team.name}}\n" +
    "                        </div>\n" +
    "                    </div>\n" +
    "                </div>\n" +
    "            </div>\n" +
    "        </div>\n" +
    "    </div>\n" +
    "</div>");
}]);

angular.module("../app/components/league/league.standings.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/league/league.standings.html",
    "<div class=\"standings-list\">\n" +
    "    <table>\n" +
    "        <tr>\n" +
    "            <th class=\"standings-team-cell\">Team</th>\n" +
    "            <th>Wins</th>\n" +
    "            <th>Losses</th>\n" +
    "            <th>Ties</th>\n" +
    "        </tr>\n" +
    "        <tr data-ng-repeat=\"standing in league.standings\">\n" +
    "           <td class=\"standings-team-cell\">{{standing.team.name}}</td>\n" +
    "           <td>{{standing.wins}}</td>\n" +
    "           <td>{{standing.losses}}</td>\n" +
    "           <td>{{standing.ties}}</td>\n" +
    "        </tr>\n" +
    "    </table>\n" +
    "</div>");
}]);

angular.module("../app/components/screen/screen.html", []).run(["$templateCache", function($templateCache) {
  $templateCache.put("../app/components/screen/screen.html",
    "<p class=\"panel\">\n" +
    "  <strong class=\"show-for-small-only\">This text is shown only on a small screen.</strong>\n" +
    "  <strong class=\"show-for-medium-up\">This text is shown on medium screens and up.</strong>\n" +
    "  <strong class=\"show-for-medium-only\">This text is shown only on a medium screen.</strong>\n" +
    "  <strong class=\"show-for-large-up\">This text is shown on large screens and up.</strong>\n" +
    "  <strong class=\"show-for-large-only\">This text is shown only on a large screen.</strong>\n" +
    "  <strong class=\"show-for-xlarge-up\">This text is shown on xlarge screens and up.</strong>\n" +
    "  <strong class=\"show-for-xlarge-only\">This text is shown only on an xlarge screen.</strong>\n" +
    "  <strong class=\"show-for-xxlarge-up\">This text is shown on xxlarge screens and up.</strong>\n" +
    "</p>");
}]);
