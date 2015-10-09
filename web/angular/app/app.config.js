angular.module('rlApp')
  .config(function(AngularyticsProvider) {
    AngularyticsProvider.setEventHandlers(['GoogleUniversal']);
  }).run(function(Angularytics) {
    Angularytics.init();
  });