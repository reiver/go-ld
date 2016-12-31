/*
Package ld provides alternatives to Go's built-in types (bool, float64, int64, string, time.Time)
that also let you express a "lack of a value".

For example, consider Go's built-in int64 versus this package's ld.Int64.

This package's ld.Int64 can express all the same values and do all the same things as
Go's built-in int64.

But this package's ld.Int64 can also express that, for example, the value for something has not
been loaded into it. Maybe not loaded from a database.

Which is a certain kind of "lack of a value". (There is another type, which we'll go over
a little later.)

Let's make that example a little more concrete. Let's consider this struct:

	type Person struct {
		ID ld.Int64

		GivenName  ld.String
		FamilyName ld.String

		HomeCountryName   ld.String
		HomeRegion        ld.String
		HomeLocality      ld.String
		HomeStreetAddress ld.String
		HomePostalCode    ld.String
	}

So, this struct is used to represent a person.

And note that we have used 2 types from this package: ld.Int64 and ld.String.

(These are this package's alternatives to Go's built-in int64 and string.)

The ID field (in the Person struct) is the primary key in our database.
(I.e., the handle that uniquely identifies this person's record in the database.)
This might have a value such as: 1234.

The GivenName field (in the Person struct) is used to store the person's given name.
(Just an FYI: some people would call a "given name" a person's "first name".)
This might have a value such as: "Fred".

The FamilyName field (in the Person struct) is used to store the person's family name.
(Just an FYI: some people would call a "family name" a person's "last name".)
This might have a value such as: "Doe".

The HomeCountryName, HomeRegion, HomeLocality, HomeStreetAddress, and HomePostalCode
fields (in the Person struct) are used to store the address of where the person lives.

So, for example, our struct could have the following values:

	ID -> 1234

	GivenName  -> "Fred"
	FamilyName -> "Doe"

	HomeCountryName   -> "Canada"
	HomeRegion        -> "BC"
	HomeLocality      -> "Vancouver"
	HomeStreetAddress -> "33 Water Street"
	HomePostalCode    -> "V6B 1R4"

Now let's say in our code, all we know about a person is their ID.

Let's say that ID is 789.

Who is this person, with ID 789?

We can get their information from database.

But let's say we only care about the person's "given name" and "family name".

In other words, we only care about the values of the GivenName and FamilyName fields, in our Person struct.

We don't care about the values of the HomeCountryName, HomeRegion, HomeLocality, HomeStreetAddress, and HomePostalCode fields.

So, when we go to our database, and tell our database "give me information for the person with ID = 789", we would also tell it
we only want it to load the GivenName and FamilyName fields.

(And not bother loading the HomeCountryName, HomeRegion, HomeLocality, HomeStreetAddress, and HomePostalCode fields.)

So, if this was a database that used SQL, instead of writing:

	SELECT *
	
	FROM persons
	
	WHERE id = 789

We would instead write:

	SELECT id
	     , given_name
	     , family_name
	
	FROM persons
	
	WHERE id = 789

(Note that I also asked the database for the "id" field too, in addition to the "given_name" and "family_name" fields.)

In case, what we would see in our struct is:

	ID -> 789

	GivenName  -> "Joe"
	FamilyName -> "Blow"

	HomeCountryName   -> StringNotLoaded
	HomeRegion        -> StringNotLoaded
	HomeLocality      -> StringNotLoaded
	HomeStreetAddress -> StringNotLoaded
	HomePostalCode    -> StringNotLoaded

What we see in the ID, GivenName, and FamilyName fields, in our Person struct, are what we expected.

(And note that person with ID = 789 has the name "Joe Blow".)

But look closely at the HomeCountryName, HomeRegion, HomeLocality, HomeStreetAddress, and HomePostalCode fields.

Those 5 fields that are used to store where the person lives are of type ld.String.

But those 5 fields don't have a string value though.

They have this other thing: StringNotLoaded

StringNotLoaded is a way of expressing a certain kind of a "lack of a value".

Because we didn't load the person's HomeCountryName, HomeRegion, HomeLocality, HomeStreetAddress, and HomePostalCode from
the database, they get the set to StringNotLoaded.


Now let's look at the other type of a "lack of a value" that some types in this package can also express.

For example, you have seen ld.String in our example. Now let us also look at ld.NullableString.

So let's say that we want to change our struct a bit. Let's say we want to add a new fields.

The way we defined our Person struct originally was:

	type Person struct {
		ID ld.Int64

		GivenName  ld.String
		FamilyName ld.String

		HomeCountryName   ld.String
		HomeRegion        ld.String
		HomeLocality      ld.String
		HomeStreetAddress ld.String
		HomePostalCode    ld.String
	}

Now, let's add a new field: AdditionalName

	type Person struct {
		ID ld.Int64

		GivenName      ld.String
		AdditionalName ld.NullableString
		FamilyName     ld.String

		HomeCountryName   ld.String
		HomeRegion        ld.String
		HomeLocality      ld.String
		HomeStreetAddress ld.String
		HomePostalCode    ld.String
	}

So, to our Person struct we have added a new field: AdditionalName. Which is used to store a
person's "additional name", if they have one.

(Just an FYI: some people would call an "additional name" a person's "middle name".)

So why did we give the AdditionalName field a type o ld.NullableString rather than ld.String?

The reason is that not everyone has an "additional name". Some people have one. Some peopel don't.

In other words, some people their "additional name" exists. And for other people their
"additional name" does NOT exist!

How do we represent this "lack of existance"?

We represent this "lack of existance" with "null".

(Note, don't confuse "null" with Go's concept of "nil", which is something pointer specific.)

This package's concept of "null" is used to represent a different kind of a "lack of a value".

It is not that a value for it exists, but we just didn't load it from the database.

It is that even if we looked at the database, the database wouldn't have a value for it.

So let's look at what we be in our Person struct for person ID = 789 (that we were looking at before).

I.e., "Joe Blow".

What if Joe Blow is one of the people who does not have an "additional name"?

Our SQL for this might change to:

	SELECT id
	     , given_name
	     , additional_name
	     , family_name
	
	FROM persons
	
	WHERE id = 789

Then what we would see is:

	ID -> 789

	GivenName      -> "Joe"
	AdditionalName -> StringNull
	FamilyName     -> "Blow"

	HomeCountryName   -> StringNotLoaded
	HomeRegion        -> StringNotLoaded
	HomeLocality      -> StringNotLoaded
	HomeStreetAddress -> StringNotLoaded
	HomePostalCode    -> StringNotLoaded

Look closely at AdditionalName.

We see what is in there is: StringNull

That StringNull expresses that we actually went to see information about AdditionalName, and what
we were told is that it does not exist.

StringNull expresses a "lack of a value".

But this is a different kind of a "lack of a value" than StringNotLoaded.
*/
package ld
