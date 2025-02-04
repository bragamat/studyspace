defmodule FreelancerRates do

  def daily_rate(hourly_rate), do: hourly_rate * 8.0

  def apply_discount(before_discount, discount) when discount == 0, do: before_discount
  def apply_discount(before_discount, discount) do
    before_discount - (before_discount * (discount/100)) |> Float.round(5)
  end

  def monthly_rate(hourly_rate, discount) do
    daily_rate(hourly_rate) * 22
    |> apply_discount(discount)
    |> Float.ceil(0)
    |> trunc
  end

  def days_in_budget(budget, hourly_rate, discount) do
    daily_rate(hourly_rate)
    |> apply_discount(discount)
    |> inverse_division(budget)
    |> Float.floor(1)
  end

  defp inverse_division(a, b), do: b / a
end
