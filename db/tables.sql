CREATE TABLE [users] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [role_id] int,
  [password] nvarchar,
  [username] nvarchar(255),
  [firstname] nvarchar(255),
  [lastname] nvarchar(255),
  [email] nvarchar(255),
  [is_private] bit
)
GO

CREATE TABLE [roles] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [rolename] nvarchar(255)
)
GO

CREATE TABLE [advertisements] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [price] int,
  [description] nvarchar,
  [title] nvarchar(255),
  [user_id] int,
  [car_id] int
)
GO

CREATE TABLE [cars] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [make] nvarchar,
  [model] nvarchar,
  [power] int,
  [price] int,
  [mileage] int,
  [emissionsCategory] int,
  [firstRegistration] datetime,
  [color] nvarchar,
  [category_id] int,
  [fuelType_id] int,
  [gearbox_id] int
)
GO

CREATE TABLE [gearboxes] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [value] nvarchar(255),
  [text] nvarchar(255)
)
GO

CREATE TABLE [categories] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [value] nvarchar(255),
  [text] nvarchar(255)
)
GO

CREATE TABLE [fueltypes] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [value] nvarchar(255),
  [text] nvarchar(255)
)
GO

CREATE TABLE [lease_offers] (
  [id] int PRIMARY KEY IDENTITY(1, 1),
  [user_id] int,
  [description] nvarchar(255),
  [interest_rate] float,
  [period_in_months] int,
  [min_value] int,
  [max_value] int,
  [first_instalment_pt_min] float,
  [first_instalment_pt_max] float,
  [required_insurance] bit
)
GO

ALTER TABLE [users] ADD FOREIGN KEY ([role_id]) REFERENCES [roles] ([id])
GO

ALTER TABLE [advertisements] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [cars] ADD FOREIGN KEY ([gearbox_id]) REFERENCES [gearboxes] ([id])
GO

ALTER TABLE [cars] ADD FOREIGN KEY ([category_id]) REFERENCES [categories] ([id])
GO

ALTER TABLE [cars] ADD FOREIGN KEY ([fuelType_id]) REFERENCES [fueltypes] ([id])
GO

ALTER TABLE [lease_offers] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO
