Clean architecture from the inner circle to the outer

use cases : don change with the change in framework or database for example, but change when the demand and explanation for the certain application changes. application specific business rules
 ==> these use cases direct entities to use their enterprise wide business rules to acheive the goals of the use case
 ==> changes in this layer wont affect the entities

Interface Adapters : convert data from a format most convenient to the use cases and entities to external agency such as db or web. contains the model view and controller of the applicaiton. models are just datastructures passed from the controllers to the use cases

code inward doesnt know anything about the DB

any other adapter necessary to convert data from some external form, to the internal form used by the use cases and entities


Frameworks and drivers : The outermost layer is composed of tools such as db, web frameworks. you don write much code here, u just glue code to communicate to the next layer. This is where all the details go. web and db for example.

source code dependencies should only point inwards. inner layers should not know anything about the outer layers.

inward => level of abstraction increases
change the ui without changing the business logic => outer most low level concrete detail

crossing boundaries

presenter == data formatter to be representible for the UI. controller => invoke necessary useCases and execute business logic

instead of writing t everytime we invoke a tast, we can use asser.New(t) for that

example:

func TestSomething(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(expected, actual, "result should be equal") //t would be passed here normally
}




