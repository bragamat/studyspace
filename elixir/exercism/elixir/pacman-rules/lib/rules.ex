defmodule Rules do

  def eat_ghost?(power_pellet_active?, _touching_ghost?)
    when power_pellet_active? == false, do: false

  def eat_ghost?(power_pellet_active?, touching_ghost?) do
    power_pellet_active? == touching_ghost?
  end

  def score?(touching_power_pellet?, touching_dot?)
    when touching_power_pellet? == false do
    touching_dot? == true
  end

  def score?(touching_power_pellet?, touching_dot?)
    when touching_power_pellet? == true do
    touching_dot? == false
  end

  def lose?(_power_pellet_active?, touching_ghost?)
    when touching_ghost? == false do

    false
  end

  def lose?(power_pellet_active?, touching_ghost?)
    when power_pellet_active? == true do
    touching_ghost? == false
  end

  def lose?(power_pellet_active?, touching_ghost?)
    when power_pellet_active? == false do

    touching_ghost? == true
  end

  def win?(_has_eaten_all_dots?, power_pellet_active?, _touching_ghost?) when
    power_pellet_active? == true do

    true
  end

  def win?(_has_eaten_all_dots?, _power_pellet_active?, touching_ghost?) when
    touching_ghost? == true do

    false
  end

  def win?(has_eaten_all_dots?, _power_pellet_active?, _touching_ghost?) when
    has_eaten_all_dots? == true do

    true
  end
end
