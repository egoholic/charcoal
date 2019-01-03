# Entities

Entities are a things that represent particular domain's business objects.

## Framework

Framework is a number of conventions. One of the most important convention is naming. Let's start with it.

### Standard Entity Methods

- **PK()** - should return a primary key of any type.
- **Attribute()** - name for read-accessors. All the state reads MUST be performed with read-accessors.
- **AssignAttribute()** - name for write-accessors. All the state changes MUST be performed with accessor methods that MUST guarantee entity-level data consistency and support valid/correct state of Entity. The main rule here is: Any moment of time the entity should be in a valid state.
- **IsPersisted()** ***-> bool*** - returs true if the entyty were persisted and false if doesn't.



## Main Blog Entities

### Account
Account represents the thing that can interact with the system.

### Content
Content represents the data that might be published.

### Publication
Publication represents the content that was published, gives it a name(PK) and stores publication date and time. **Publication is an ACT of publishing of some content**. If there is a Publication for a Content, than the contact was published. You can also interpret Publication as an event. Because event ~ fact. And,  because of a publication is a FACT, it SHOULD NOT be mutated! The only mutable field of a publication is ``views`` which stores the number of publication's views.
