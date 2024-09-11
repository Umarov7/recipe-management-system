INSERT INTO users (id, username, email, password_hash, role) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'johndoe', 'johndoe@example.com', 'hashed_password_1', 'chef'),
('550e8400-e29b-41d4-a716-446655440001', 'janedoe', 'janedoe@example.com', 'hashed_password_2', 'admin'),
('550e8400-e29b-41d4-a716-446655440002', 'alice', 'alice@example.com', 'hashed_password_3', 'user'),
('550e8400-e29b-41d4-a716-446655440003', 'bob', 'bob@example.com', 'hashed_password_4', 'user'),
('550e8400-e29b-41d4-a716-446655440004', 'charlie', 'charlie@example.com', 'hashed_password_5', 'admin'),
('550e8400-e29b-41d4-a716-446655440005', 'david', 'david@example.com', 'hashed_password_6', 'chef'),
('550e8400-e29b-41d4-a716-446655440006', 'eve', 'eve@example.com', 'hashed_password_7', 'user'),
('550e8400-e29b-41d4-a716-446655440007', 'frank', 'frank@example.com', 'hashed_password_8', 'admin'),
('550e8400-e29b-41d4-a716-446655440008', 'grace', 'grace@example.com', 'hashed_password_9', 'chef'),
('550e8400-e29b-41d4-a716-446655440009', 'heidi', 'heidi@example.com', 'hashed_password_10', 'user');

INSERT INTO recipes (id, user_id, category, title, description, instructions, prep_time, cook_time, servings) VALUES
('550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440000', 'Breakfast', 'Classic Pancakes', 'Fluffy pancakes perfect for breakfast.', '1. Mix dry ingredients. 2. Add wet ingredients. 3. Cook on griddle.', 10, 15, 4),
('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440005', 'Lunch', 'Grilled Chicken Salad', 'A healthy and tasty grilled chicken salad.', '1. Grill chicken. 2. Toss salad ingredients. 3. Add grilled chicken.', 20, 10, 2),
('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440008', 'Dinner', 'Spaghetti Bolognese', 'A rich and savory spaghetti bolognese.', '1. Cook spaghetti. 2. Prepare sauce with ground beef. 3. Mix and serve.', 15, 45, 4),
('550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440008', 'Dinner', 'Beef Stroganoff', 'Tender beef in a creamy mushroom sauce.', '1. Brown beef. 2. Cook mushrooms and onions. 3. Combine and serve.', 20, 30, 4),
('550e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440005', 'Dessert', 'Chocolate Brownies', 'Rich and gooey chocolate brownies.', '1. Mix ingredients. 2. Bake in oven. 3. Cool and cut into squares.', 15, 25, 12),
('550e8400-e29b-41d4-a716-446655440015', '550e8400-e29b-41d4-a716-446655440000', 'Snacks', 'Homemade Granola Bars', 'Crunchy granola bars with nuts and dried fruit.', '1. Mix ingredients. 2. Bake and cool. 3. Cut into bars.', 10, 20, 10),
('550e8400-e29b-41d4-a716-446655440016', '550e8400-e29b-41d4-a716-446655440005', 'Breakfast', 'Avocado Toast', 'Simple and healthy avocado toast.', '1. Toast bread. 2. Mash avocado. 3. Spread avocado on toast.', 5, 0, 1),
('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440008', 'Lunch', 'Chicken Caesar Wrap', 'A wrap filled with chicken, lettuce, and Caesar dressing.', '1. Cook chicken. 2. Assemble wrap with lettuce and dressing. 3. Roll and serve.', 15, 5, 1),
('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440008', 'Dinner', 'Vegetable Stir-Fry', 'Quick and easy vegetable stir-fry with a savory sauce.', '1. Stir-fry vegetables. 2. Add sauce. 3. Serve with rice.', 10, 10, 4),
('550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440005', 'Dessert', 'Apple Crisp', 'Warm apple crisp with a buttery topping.', '1. Prepare apples. 2. Make topping. 3. Bake and serve warm.', 20, 40, 6),
('550e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440000', 'Snack', 'Hummus and Veggies', 'A healthy snack of hummus with fresh vegetables.', '1. Prepare hummus. 2. Slice veggies. 3. Serve together.', 15, 0, 2);

INSERT INTO ingredients (recipe_id, name, quantity, unit) VALUES
-- Ingredients for Classic Pancakes
('550e8400-e29b-41d4-a716-446655440010', 'Flour', 1.50, 'cups'),
('550e8400-e29b-41d4-a716-446655440010', 'Sugar', 2, 'tablespoons'),
('550e8400-e29b-41d4-a716-446655440010', 'Baking Powder', 2, 'teaspoons'),
('550e8400-e29b-41d4-a716-446655440010', 'Salt', 0.50, 'teaspoon'),
('550e8400-e29b-41d4-a716-446655440010', 'Milk', 1.25, 'cups'),
('550e8400-e29b-41d4-a716-446655440010', 'Egg', 1, 'large'),
('550e8400-e29b-41d4-a716-446655440010', 'Butter', 2, 'tablespoons'),
-- Ingredients for Grilled Chicken Salad
('550e8400-e29b-41d4-a716-446655440011', 'Chicken Breast', 2, 'pieces'),
('550e8400-e29b-41d4-a716-446655440011', 'Mixed Greens', 4, 'cups'),
('550e8400-e29b-41d4-a716-446655440011', 'Cherry Tomatoes', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440011', 'Cucumber', 1, 'medium'),
('550e8400-e29b-41d4-a716-446655440011', 'Caesar Dressing', 0.25, 'cup'),
-- Ingredients for Spaghetti Bolognese
('550e8400-e29b-41d4-a716-446655440012', 'Spaghetti', 8, 'ounces'),
('550e8400-e29b-41d4-a716-446655440012', 'Ground Beef', 1, 'pound'),
('550e8400-e29b-41d4-a716-446655440012', 'Tomato Sauce', 2, 'cups'),
('550e8400-e29b-41d4-a716-446655440012', 'Onion', 1, 'large'),
('550e8400-e29b-41d4-a716-446655440012', 'Garlic Cloves', 2, 'cloves'),
-- Ingredients for Beef Stroganoff
('550e8400-e29b-41d4-a716-446655440013', 'Beef Sirloin', 1, 'pound'),
('550e8400-e29b-41d4-a716-446655440013', 'Mushrooms', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440013', 'Onion', 1, 'medium'),
('550e8400-e29b-41d4-a716-446655440013', 'Sour Cream', 0.50, 'cup'),
('550e8400-e29b-41d4-a716-446655440013', 'Beef Broth', 1, 'cup'),
-- Ingredients for Chocolate Brownies
('550e8400-e29b-41d4-a716-446655440014', 'Butter', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440014', 'Sugar', 2, 'cups'),
('550e8400-e29b-41d4-a716-446655440014', 'Cocoa Powder', 0.75, 'cup'),
('550e8400-e29b-41d4-a716-446655440014', 'Eggs', 4, 'large'),
('550e8400-e29b-41d4-a716-446655440014', 'Flour', 1, 'cup'),
-- Ingredients for Homemade Granola Bars
('550e8400-e29b-41d4-a716-446655440015', 'Oats', 2, 'cups'),
('550e8400-e29b-41d4-a716-446655440015', 'Honey', 0.75, 'cup'),
('550e8400-e29b-41d4-a716-446655440015', 'Almonds', 0.50, 'cup'),
('550e8400-e29b-41d4-a716-446655440015', 'Dried Fruit', 0.50, 'cup'),
-- Ingredients for Avocado Toast
('550e8400-e29b-41d4-a716-446655440016', 'Bread', 1, 'slice'),
('550e8400-e29b-41d4-a716-446655440016', 'Avocado', 1, 'medium'),
-- Ingredients for Chicken Caesar Wrap
('550e8400-e29b-41d4-a716-446655440017', 'Chicken Breast', 1, 'piece'),
('550e8400-e29b-41d4-a716-446655440017', 'Tortilla Wrap', 1, 'large'),
('550e8400-e29b-41d4-a716-446655440017', 'Romaine Lettuce', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440017', 'Caesar Dressing', 2, 'tablespoons'),
-- Ingredients for Vegetable Stir-Fry
('550e8400-e29b-41d4-a716-446655440018', 'Mixed Vegetables', 4, 'cups'),
('550e8400-e29b-41d4-a716-446655440018', 'Stir-Fry Sauce', 0.25, 'cup'),
('550e8400-e29b-41d4-a716-446655440018', 'Rice', 2, 'cups'),
-- Ingredients for Apple Crisp
('550e8400-e29b-41d4-a716-446655440019', 'Apples', 6, 'medium'),
('550e8400-e29b-41d4-a716-446655440019', 'Brown Sugar', 0.75, 'cup'),
('550e8400-e29b-41d4-a716-446655440019', 'Oats', 0.50, 'cup'),
('550e8400-e29b-41d4-a716-446655440019', 'Butter', 0.50, 'cup'),
-- Ingredients for Hummus and Veggies
('550e8400-e29b-41d4-a716-446655440020', 'Hummus', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440020', 'Carrot Sticks', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440020', 'Cucumber Slices', 1, 'cup'),
('550e8400-e29b-41d4-a716-446655440020', 'Bell Pepper Strips', 1, 'cup');