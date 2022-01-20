/**
 * Hippo.Web
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.HippoWeb);
  }
}(this, function(expect, HippoWeb) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new HippoWeb.EnvironmentVariableApi();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('EnvironmentVariableApi', function() {
    describe('apiEnvironmentvariableGet', function() {
      it('should call apiEnvironmentvariableGet successfully', function(done) {
        //uncomment below and update the code to test apiEnvironmentvariableGet
        //instance.apiEnvironmentvariableGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('apiEnvironmentvariableIdDelete', function() {
      it('should call apiEnvironmentvariableIdDelete successfully', function(done) {
        //uncomment below and update the code to test apiEnvironmentvariableIdDelete
        //instance.apiEnvironmentvariableIdDelete(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('apiEnvironmentvariableIdGet', function() {
      it('should call apiEnvironmentvariableIdGet successfully', function(done) {
        //uncomment below and update the code to test apiEnvironmentvariableIdGet
        //instance.apiEnvironmentvariableIdGet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('apiEnvironmentvariablePost', function() {
      it('should call apiEnvironmentvariablePost successfully', function(done) {
        //uncomment below and update the code to test apiEnvironmentvariablePost
        //instance.apiEnvironmentvariablePost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
