package templates

var Config = `---
# input.path specifies the
# path of the file containing
# the data desired to be processed
# throughout the pipeline
input:
    path: "inputPath"
---
# output.path specifies where
# the end result of the pipeline
# will be written to
output:
    path: "outputPath"
---
# rename module(s) to the whatever
# you wish and provide the path to
# each as its value. these module
# names can then be used to denote
# the desired module to be used
# for each step below in the "steps"
# section instead of providing a path
modules:
    moduleOneName: moduleOnePath
    moduleTwoName: moduleTwoPath
    moduleThreeName: moduleThreePath
---
# global allows you to set variables
# that will be available to all steps
global:
    vars:
        globalVarOne: "globalVarValueOne"
        globalVarTwo: "globalVarValueTwo"
---
# rename step(s) to the function name to be called
# the module the step uses
#
# to set a module for a step, simply use the
# name of the desired module from the modules
# section above
#
# steps can have their own variables as well, and if
# a step has a variable with the same name as a global
# variable, then the step-specific variable will take
# precedence
steps:
  step1:
    module: moduleOneName
    vars:
      hello: world
      goodbye: all
  step2:
    module: moduleTwoName
    vars:
      hello: again
      goodbye: final
  step3:
    # if you want to run multiple steps concurrently,
    # just specify the step is concurrent and list
    # all steps to be run together under it
    concurrent:
      step4:
        module: moduleOneName
        vars:
          yes: true
          no: false
      step5:
        module: moduleThreeName
        # if a step requires making HTTP calls, then
        # you can provide a list of hosts that the
        # step will be allows to make the requests from
        # if you want to allow any hosts to be available
        # to make requests from, you can simple provide
        # the wildcard "*" as shown below
        hosts:
          - "*"`
