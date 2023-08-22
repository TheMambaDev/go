## Go structs
In Go, a struct is a composite data type that groups together zero or more fields with different 
data types under a single named type. It's similar to a class in object-oriented programming languages, but without methods. 
Structs are commonly used to represent complex data structures and are a fundamental part of Go's type system.

composition is often achieved through struct embedding, where you embed one struct type within another. This allows the fields and methods of the embedded struct to become part of the outer struct. This concept is similar to inheritance in class-based languages, but it's more flexible and doesn't have the same complexities.

## Advantages of Composition in Go

Composition is a powerful concept in object-oriented programming that offers several advantages:

1. **Code Reusability:** Composition allows you to create reusable components and combine them to construct more intricate structures. This approach fosters the development of modular and maintainable codebases.

2. **Flexibility:** With composition, you have the freedom to mix and match components, enabling the creation of diverse configurations. In contrast to inheritance, which locks you into a single hierarchy, composition empowers you to modify and adapt components easily.

3. **Encapsulation:** Composition promotes encapsulation by enabling components to have their own private fields and methods. The outer struct can expose only the necessary functionality, keeping implementation details hidden.

4. **Avoiding Diamond Problem:** Composition sidesteps the "diamond problem" that can emerge in languages with multiple inheritance. This problem involves ambiguities caused by multiple inheritance paths, which composition elegantly circumvents.

5. **Less Coupling:** Composition encourages loose coupling between components, enhancing the resilience of your codebase to changes. This means that modifications to one component are less likely to impact others.

6. **Testability:** Components can be tested independently when using composition, leading to more straightforward unit testing. Isolating the testing of individual components enhances the reliability of your tests.

7. **Clearer Hierarchy:** Each struct in a composition has explicit fields representing its components. This clear structuring improves the visibility of relationships between different entities within your codebase.

