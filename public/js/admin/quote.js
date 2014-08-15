define(['alertify'],function(alertify){
	$(document).on('click','.delete',function(e){
		var that = this;
		alertify.confirm('Are you sure you want to delete this quote?',function(e){
			var id = $(that).data('id');
			$.ajax({
				url: '/admin/quotes/'+id,
				type: 'DELETE',
				success:function(data, status, xhr){
					$(that).closest('tr').remove();
					alertify.success('Quote removed');
				},
				error: function(xhr, status, err){
					alertify.error(xhr.responseText);
				}
			});
		});
		return false;
	}).on('submit','.heading-form',function(e){
		var heading = $("#heading").val();
		$.ajax({
				url: '/admin/quotes/heading',
				type: 'POST',
				data:{
					heading:heading
				},
				success:function(data, status, xhr){
					alertify.success('Heading saved');
				},
				error: function(xhr, status, err){
					alertify.error(xhr.responseText);
				}
			});
		return false;
	});
});